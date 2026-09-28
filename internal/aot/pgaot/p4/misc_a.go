package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AdjustMicroseconds(m *base.Module, l0 int64, l1 float64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v42 int64
	_ = v42
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v69 float64
	_ = v69
	var v70 int64
	_ = v70
	var v72 float64
	_ = v72
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v95 int32
	_ = v95
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int64(63)
	v21 = int64(32)
	v22 = int64(base.Ui64(l2) >> (uint(v21) % 64))
	v24 = int64(base.Ui64(l0) >> (uint(v21) % 64))
	v27 = int64(4294967295)
	v28 = l2 & v27
	v30 = l0 & v27
	v31 = v28 * v30
	v35 = int64(base.Ui64(v31)>>(uint(v21)%64)) + v28*v24
	v42 = v30*v22 + v35&v27
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = l0*(l2>>(uint(v13)%64)) + l0>>(uint(v13)%64)*l2 + v22*v24 + int64(base.Ui64(v35)>>(uint(v21)%64)) + int64(base.Ui64(v42)>>(uint(v21)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v31&v27 | v42<<(uint(v21)%64)
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v53 != v54>>(uint(int64(63))%64) {
		v95 = v5
	} else {
		v58 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
		v59 = v58 + v54
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v59
		if base.B2i32(v54 < int64(0))^base.B2i32(v59 < v58) != 0 {
			v95 = v5
		} else {
			if base.F64_eq(l1, float64(0)) != 0 {
				v95 = int32(1)
			} else {
				v69 = base.F64_mul(l1, base.F64_convert_i64_u(l2))
				v70 = base.I64_trunc_sat_f64_s(v69)
				v72 = base.F64_sub(v69, base.F64_convert_i64_s(v70))
				if base.F64_gt(v72, float64(0.5)) != 0 {
					v83 = v70 + int64(1)
				} else {
					if base.F64_lt(v72, float64(-0.5)) == int32(0) {
						v83 = v70
					} else {
						v83 = v70 - int64(1)
					}
				}
				v84 = v59 + v83
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v84
				v95 = base.B2i32(base.B2i32(v83 < int64(0))^base.B2i32(v84 < v59) == int32(0))
			}
		}
	}
	m.G0 = v11 + int32(16)
	return v95
}
func F_AdvanceNextFullTransactionIdPastXid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	v4 = int32(3)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = base.I32_wrap_i64(v8)
	if base.B2i32(base.Ui32(l0) < base.Ui32(v4))|base.B2i32(base.Ui32(v9) < base.Ui32(v4)) == int32(0) {
		if int32(0) <= l0-v9 {
			v20 = int64(base.Ui64(v8) >> (uint(int64(32)) % 64))
			v25 = int32(3)
			v27 = l0 + int32(1)
			if base.Ui32(v27) <= base.Ui32(v25) {
				v30 = v25
			} else {
				v30 = v27
			}
			if base.Ui32(v30) < base.Ui32(v9) {
				v32 = (v20 + int64(1)) & int64(4294967295)
			} else {
				v32 = v20
			}
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
			v38 = F_LWLockAcquire(m, v34+int32(384), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = base.I64_extend_i32_u(v30) | v32<<(uint(int64(32))%64)
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
				F_LWLockRelease(m, v48+int32(384))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	} else {
		if base.Ui32(l0) < base.Ui32(v9) {
			return
		} else {
			v20 = int64(base.Ui64(v8) >> (uint(int64(32)) % 64))
			v25 = int32(3)
			v27 = l0 + int32(1)
			if base.Ui32(v27) <= base.Ui32(v25) {
				v30 = v25
			} else {
				v30 = v27
			}
			if base.Ui32(v30) < base.Ui32(v9) {
				v32 = (v20 + int64(1)) & int64(4294967295)
			} else {
				v32 = v20
			}
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
			v38 = F_LWLockAcquire(m, v34+int32(384), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = base.I64_extend_i32_u(v30) | v32<<(uint(int64(32))%64)
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
				F_LWLockRelease(m, v48+int32(384))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_AioShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemAttach[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		m.T0[v5].(func(*base.Module, int32))(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_AllocSetDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v5 < int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v8 != 0 {
			v12 = v8
			for {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				if v12 != l0+int32(112) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + (v12 - v18)
					F_emscripten_builtin_free(m, v12)
					mBase = m.M
				} else {
				}
				if v15 != 0 {
					v12 = v15
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		F_emscripten_builtin_free(m, l0)
		mBase = m.M
		return
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v28 == int32(0) {
			F_MemoryContextResetOnly(m, l0)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = v5 << (uint(int32(3)) % 32)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
				if v37 < int32(100) {
					v58 = v37
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1])))
					if v40 == int32(0) {
						v58 = v37
					} else {
						v44 = v40
						for {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1]))) = v47
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
							*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0]))) = v49 - int32(1)
							F_emscripten_builtin_free(m, v44)
							mBase = m.M
							if v47 != 0 {
								v44 = v47
								continue
							} else {
								break
							}
							break
						}
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
						v58 = v54
					}
				}
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1])))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0]))) = v58 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1]))) = l0
				return
			}
		} else {
			v34 = v5 << (uint(int32(3)) % 32)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
			if v37 < int32(100) {
				v58 = v37
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1])))
				if v40 == int32(0) {
					v58 = v37
				} else {
					v44 = v40
					for {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1]))) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0]))) = v49 - int32(1)
						F_emscripten_builtin_free(m, v44)
						mBase = m.M
						if v47 != 0 {
							v44 = v47
							continue
						} else {
							break
						}
						break
					}
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0])))
					v58 = v54
				}
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1])))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[0]))) = v58 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_AllocSetDelete[1]))) = l0
			return
		}
	}
}
func F_AllocSetGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v5 = l0 - int32(8)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6&int64(16) != int64(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(16))))
		return v13 - v5
	} else {
		v16 = int32(8)
		return v16<<(uint(base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(5))%64))))%32) + v16
	}
}
func F_AlterForeignServerOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v14 = v12 + v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	if l2 != v15 {
		v17 = F_superuser(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(65536)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_extend_i32_u(l2)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v61 = F_heap_getattr_8(m, l1, v58, v10+int32(15))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v63 == int32(0) {
						v67 = F_pg_detoast_datum(m, base.I32_wrap_i64(v61))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
							v70 = F_aclnewowner(m, v67, v69, l2)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v72)
								*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = base.I64_extend_i32_u(v70)
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
										F_changeDependencyOnOwner(m, int32(1417), v91, l2)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
											if v98 != 0 {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
												v101 = int32(0)
												F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													m.G0 = v10 + int32(96)
													return
												}
											} else {
												m.G0 = v10 + int32(96)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
								F_changeDependencyOnOwner(m, int32(1417), v91, l2)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
									if v98 != 0 {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
										v101 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											m.G0 = v10 + int32(96)
											return
										}
									} else {
										m.G0 = v10 + int32(96)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
				v23 = F_object_ownercheck(m, int32(1417), v20, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 == int32(0) {
						F_aclcheck_error(m, int32(2), int32(17), v14+int32(4))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
							F_check_can_set_role(m, v34, l2)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
								v40 = F_object_aclcheck(m, int32(2328), v38, l2, int64(256))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									if v40 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(65536)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_extend_i32_u(l2)
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v61 = F_heap_getattr_8(m, l1, v58, v10+int32(15))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v63 == int32(0) {
												v67 = F_pg_detoast_datum(m, base.I32_wrap_i64(v61))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
													v70 = F_aclnewowner(m, v67, v69, l2)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return
													} else {
														v72 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v72)
														*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = base.I64_extend_i32_u(v70)
														v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
														v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																F_changeDependencyOnOwner(m, int32(1417), v91, l2)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																	if v98 != 0 {
																		v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																		v101 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(96)
																			return
																		}
																	} else {
																		m.G0 = v10 + int32(96)
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
														F_changeDependencyOnOwner(m, int32(1417), v91, l2)
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
															if v98 != 0 {
																v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																v101 = int32(0)
																F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(96)
																	return
																}
															} else {
																m.G0 = v10 + int32(96)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
										v46 = F_GetForeignDataWrapper(m, v45)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
											F_aclcheck_error(m, v40, int32(16), v48)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(65536)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_extend_i32_u(l2)
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v61 = F_heap_getattr_8(m, l1, v58, v10+int32(15))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v63 == int32(0) {
														v67 = F_pg_detoast_datum(m, base.I32_wrap_i64(v61))
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
															v70 = F_aclnewowner(m, v67, v69, l2)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return
															} else {
																v72 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v72)
																*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = base.I64_extend_i32_u(v70)
																v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
																v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return
																	} else {
																		v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																		F_changeDependencyOnOwner(m, int32(1417), v91, l2)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return
																		} else {
																			v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																			if v98 != 0 {
																				v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																				v101 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return
																				} else {
																					m.G0 = v10 + int32(96)
																					return
																				}
																			} else {
																				m.G0 = v10 + int32(96)
																				return
																			}
																		}
																	}
																}
															}
														}
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
														v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																F_changeDependencyOnOwner(m, int32(1417), v91, l2)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																	if v98 != 0 {
																		v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																		v101 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			m.G0 = v10 + int32(96)
																			return
																		}
																	} else {
																		m.G0 = v10 + int32(96)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
						F_check_can_set_role(m, v34, l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
							v40 = F_object_aclcheck(m, int32(2328), v38, l2, int64(256))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								if v40 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(65536)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_extend_i32_u(l2)
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v61 = F_heap_getattr_8(m, l1, v58, v10+int32(15))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v63 == int32(0) {
											v67 = F_pg_detoast_datum(m, base.I32_wrap_i64(v61))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
												v70 = F_aclnewowner(m, v67, v69, l2)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													v72 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v72)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = base.I64_extend_i32_u(v70)
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															F_changeDependencyOnOwner(m, int32(1417), v91, l2)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																if v98 != 0 {
																	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	v101 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(96)
																		return
																	}
																} else {
																	m.G0 = v10 + int32(96)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
													F_changeDependencyOnOwner(m, int32(1417), v91, l2)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
														if v98 != 0 {
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															v101 = int32(0)
															F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																m.G0 = v10 + int32(96)
																return
															}
														} else {
															m.G0 = v10 + int32(96)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
									v46 = F_GetForeignDataWrapper(m, v45)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
										F_aclcheck_error(m, v40, int32(16), v48)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(65536)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_extend_i32_u(l2)
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v61 = F_heap_getattr_8(m, l1, v58, v10+int32(15))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												if v63 == int32(0) {
													v67 = F_pg_detoast_datum(m, base.I32_wrap_i64(v61))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
														v70 = F_aclnewowner(m, v67, v69, l2)
														mBase = m.M
														v71 = m.ExcPending
														if v71 != 0 {
															return
														} else {
															v72 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v72)
															*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = base.I64_extend_i32_u(v70)
															v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
															v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	F_changeDependencyOnOwner(m, int32(1417), v91, l2)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																		if v98 != 0 {
																			v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																			v101 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return
																			} else {
																				m.G0 = v10 + int32(96)
																				return
																			}
																		} else {
																			m.G0 = v10 + int32(96)
																			return
																		}
																	}
																}
															}
														}
													}
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v84 = F_heap_modify_tuple(m, l1, v77, v10+int32(32), v10+int32(24), v10+int32(16))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l0, v84+int32(4), v84)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
															F_changeDependencyOnOwner(m, int32(1417), v91, l2)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																if v98 != 0 {
																	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
																	v101 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
																		return
																	} else {
																		m.G0 = v10 + int32(96)
																		return
																	}
																} else {
																	m.G0 = v10 + int32(96)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v98 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
		if v98 != 0 {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v101 = int32(0)
			F_RunObjectPostAlterHook(m, int32(1417), v100, v101, v101, v101)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return
			} else {
				m.G0 = v10 + int32(96)
				return
			}
		} else {
			m.G0 = v10 + int32(96)
			return
		}
	}
}
func F_AppendJumble16(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v376 = v11
	} else {
		v12 = int32(4)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
			v23 = v14
			v25 = v12
			v27 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v23) {
					v30 = int32(1024)
					v37 = int32(-1636607408)
					if v13&int32(3) != 0 {
						v83 = v13
						v84 = v30
						v86 = v37
						v87 = v37
						v88 = v37
						for {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
							v91 = v90 + v87
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
							v95 = v94 + v88
							v97 = int32(4)
							v99 = v92 + v86 - v95 ^ base.I32_rotl(v95, v97)
							v103 = v91 - v99 ^ base.I32_rotl(v99, int32(6))
							v104 = v95 + v91
							v105 = v99 + v104
							v106 = v103 + v105
							v110 = v104 - v103 ^ base.I32_rotl(v103, int32(8))
							v114 = v105 - v110 ^ base.I32_rotl(v110, int32(16))
							v118 = v106 - v114 ^ base.I32_rotl(v114, int32(19))
							v119 = v110 + v106
							v120 = v114 + v119
							v121 = v118 + v120
							v125 = v119 - v118 ^ base.I32_rotl(v118, v97)
							v126 = int32(12)
							v127 = v83 + v126
							v129 = v84 - v126
							if base.Ui32(int32(11)) < base.Ui32(v129) {
								v83 = v127
								v84 = v129
								v86 = v120
								v87 = v121
								v88 = v125
								continue
							} else {
								break
							}
							break
						}
						switch v129 - int32(1) {
						case 0:
							v302 = v120
							v303 = v121
							v304 = v125
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 1:
							v295 = v120
							v296 = v121
							v297 = v125
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 2:
							v288 = v120
							v289 = v121
							v290 = v125
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 3:
							v282 = v121
							v283 = v125
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 4:
							v278 = v121
							v279 = v125
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 5:
							v272 = v121
							v273 = v125
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 6:
							v266 = v121
							v267 = v125
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 7:
							v261 = v125
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 8:
							v256 = v125
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 9:
							v251 = v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						case 10:
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
							v251 = v247<<(uint(int32(24))%32) + v125
							v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
							v256 = v252<<(uint(int32(16))%32) + v251
							v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
							v261 = v257<<(uint(int32(8))%32) + v256
							v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
							v266 = v262<<(uint(int32(24))%32) + v121
							v267 = v261
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
							v272 = v268<<(uint(int32(16))%32) + v266
							v273 = v267
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
							v278 = v274<<(uint(int32(8))%32) + v272
							v279 = v273
							v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
							v282 = v278 + v280
							v283 = v279
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
							v288 = v284<<(uint(int32(24))%32) + v120
							v289 = v282
							v290 = v283
							v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
							v295 = v291<<(uint(int32(16))%32) + v288
							v296 = v289
							v297 = v290
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
							v302 = v298<<(uint(int32(8))%32) + v295
							v303 = v296
							v304 = v297
							v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
							v310 = v302 + v305
							v311 = v303
							v312 = v304
						default:
							v310 = v120
							v311 = v121
							v312 = v125
						}
					} else {
						v143 = v13
						v144 = v30
						v146 = v37
						v147 = v37
						v148 = v37
						for {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
							v151 = v150 + v147
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
							v155 = v154 + v148
							v157 = int32(4)
							v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
							v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
							v164 = v155 + v151
							v165 = v159 + v164
							v166 = v163 + v165
							v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
							v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
							v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
							v179 = v170 + v166
							v180 = v174 + v179
							v181 = v178 + v180
							v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
							v186 = int32(12)
							v187 = v143 + v186
							v189 = v144 - v186
							if base.Ui32(int32(11)) < base.Ui32(v189) {
								v143 = v187
								v144 = v189
								v146 = v180
								v147 = v181
								v148 = v185
								continue
							} else {
								break
							}
							break
						}
						switch v189 - int32(1) {
						case 0:
							v244 = v180
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 1:
							v239 = v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 2:
							v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
							v239 = v235<<(uint(int32(16))%32) + v180
							v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
							v244 = v240<<(uint(int32(8))%32) + v239
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
							v310 = v244 + v245
							v311 = v181
							v312 = v185
						case 3:
							v232 = v181
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 4:
							v229 = v181
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 5:
							v224 = v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 6:
							v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
							v224 = v220<<(uint(int32(16))%32) + v181
							v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
							v229 = v225<<(uint(int32(8))%32) + v224
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
							v232 = v229 + v230
							v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v310 = v233 + v180
							v311 = v232
							v312 = v185
						case 7:
							v215 = v185
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 8:
							v210 = v185
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 9:
							v205 = v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						case 10:
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
							v205 = v201<<(uint(int32(24))%32) + v185
							v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
							v210 = v206<<(uint(int32(16))%32) + v205
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
							v215 = v211<<(uint(int32(8))%32) + v210
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
							v310 = v216 + v180
							v311 = v218 + v181
							v312 = v215
						default:
							v310 = v180
							v311 = v181
							v312 = v185
						}
					}
					v315 = int32(14)
					v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
					v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
					v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
					v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
					v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
					v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
					*(*int64)(unsafe.Add(mBase, uint32(v13))) = base.I64_extend_i32_u(v337)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24)))
					v349 = int32(8)
				} else {
					v349 = v23
				}
				v351 = int32(1024) - v349
				if base.Ui32(v25) < base.Ui32(v351) {
					v353 = v25
				} else {
					v353 = v351
				}
				if v353 != 0 {
					base.MemoryCopy(m, v349+v13, v27, v353)
				} else {
				}
				v357 = v349 + v353
				v358 = v25 - v353
				if v358 != 0 {
					v23 = v357
					v25 = v358
					v27 = v353 + v27
					continue
				} else {
					break
				}
				break
			}
			v366 = v357
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
			v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v366 = v361 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
		v376 = v366
	}
	v381 = int32(2)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v376-int32(1023)) < base.Ui32(v381) {
		v388 = l1
		v389 = v376
		v391 = v381
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v389) {
				v396 = int32(1024)
				v403 = int32(-1636607408)
				if v382&int32(3) != 0 {
					v449 = v382
					v450 = v396
					v452 = v403
					v453 = v403
					v454 = v403
					for {
						v456 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
						v457 = v456 + v453
						v458 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
						v460 = *(*int32)(unsafe.Add(mBase, uint32(v449)+8))
						v461 = v460 + v454
						v463 = int32(4)
						v465 = v458 + v452 - v461 ^ base.I32_rotl(v461, v463)
						v469 = v457 - v465 ^ base.I32_rotl(v465, int32(6))
						v470 = v461 + v457
						v471 = v465 + v470
						v472 = v469 + v471
						v476 = v470 - v469 ^ base.I32_rotl(v469, int32(8))
						v480 = v471 - v476 ^ base.I32_rotl(v476, int32(16))
						v484 = v472 - v480 ^ base.I32_rotl(v480, int32(19))
						v485 = v476 + v472
						v486 = v480 + v485
						v487 = v484 + v486
						v491 = v485 - v484 ^ base.I32_rotl(v484, v463)
						v492 = int32(12)
						v493 = v449 + v492
						v495 = v450 - v492
						if base.Ui32(int32(11)) < base.Ui32(v495) {
							v449 = v493
							v450 = v495
							v452 = v486
							v453 = v487
							v454 = v491
							continue
						} else {
							break
						}
						break
					}
					switch v495 - int32(1) {
					case 0:
						v668 = v486
						v669 = v487
						v670 = v491
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 1:
						v661 = v486
						v662 = v487
						v663 = v491
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 2:
						v654 = v486
						v655 = v487
						v656 = v491
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 3:
						v648 = v487
						v649 = v491
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 4:
						v644 = v487
						v645 = v491
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 5:
						v638 = v487
						v639 = v491
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 6:
						v632 = v487
						v633 = v491
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 7:
						v627 = v491
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 8:
						v622 = v491
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 9:
						v617 = v491
						v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+9)))
						v622 = v618<<(uint(int32(16))%32) + v617
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					case 10:
						v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+10)))
						v617 = v613<<(uint(int32(24))%32) + v491
						v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+9)))
						v622 = v618<<(uint(int32(16))%32) + v617
						v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+8)))
						v627 = v623<<(uint(int32(8))%32) + v622
						v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+7)))
						v632 = v628<<(uint(int32(24))%32) + v487
						v633 = v627
						v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+6)))
						v638 = v634<<(uint(int32(16))%32) + v632
						v639 = v633
						v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+5)))
						v644 = v640<<(uint(int32(8))%32) + v638
						v645 = v639
						v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+4)))
						v648 = v644 + v646
						v649 = v645
						v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+3)))
						v654 = v650<<(uint(int32(24))%32) + v486
						v655 = v648
						v656 = v649
						v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+2)))
						v661 = v657<<(uint(int32(16))%32) + v654
						v662 = v655
						v663 = v656
						v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
						v668 = v664<<(uint(int32(8))%32) + v661
						v669 = v662
						v670 = v663
						v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
						v676 = v668 + v671
						v677 = v669
						v678 = v670
					default:
						v676 = v486
						v677 = v487
						v678 = v491
					}
				} else {
					v509 = v382
					v510 = v396
					v512 = v403
					v513 = v403
					v514 = v403
					for {
						v516 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
						v517 = v516 + v513
						v518 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
						v520 = *(*int32)(unsafe.Add(mBase, uint32(v509)+8))
						v521 = v520 + v514
						v523 = int32(4)
						v525 = v518 + v512 - v521 ^ base.I32_rotl(v521, v523)
						v529 = v517 - v525 ^ base.I32_rotl(v525, int32(6))
						v530 = v521 + v517
						v531 = v525 + v530
						v532 = v529 + v531
						v536 = v530 - v529 ^ base.I32_rotl(v529, int32(8))
						v540 = v531 - v536 ^ base.I32_rotl(v536, int32(16))
						v544 = v532 - v540 ^ base.I32_rotl(v540, int32(19))
						v545 = v536 + v532
						v546 = v540 + v545
						v547 = v544 + v546
						v551 = v545 - v544 ^ base.I32_rotl(v544, v523)
						v552 = int32(12)
						v553 = v509 + v552
						v555 = v510 - v552
						if base.Ui32(int32(11)) < base.Ui32(v555) {
							v509 = v553
							v510 = v555
							v512 = v546
							v513 = v547
							v514 = v551
							continue
						} else {
							break
						}
						break
					}
					switch v555 - int32(1) {
					case 0:
						v610 = v546
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 1:
						v605 = v546
						v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
						v610 = v606<<(uint(int32(8))%32) + v605
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 2:
						v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+2)))
						v605 = v601<<(uint(int32(16))%32) + v546
						v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
						v610 = v606<<(uint(int32(8))%32) + v605
						v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
						v676 = v610 + v611
						v677 = v547
						v678 = v551
					case 3:
						v598 = v547
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 4:
						v595 = v547
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 5:
						v590 = v547
						v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+5)))
						v595 = v591<<(uint(int32(8))%32) + v590
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 6:
						v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+6)))
						v590 = v586<<(uint(int32(16))%32) + v547
						v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+5)))
						v595 = v591<<(uint(int32(8))%32) + v590
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+4)))
						v598 = v595 + v596
						v599 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v676 = v599 + v546
						v677 = v598
						v678 = v551
					case 7:
						v581 = v551
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 8:
						v576 = v551
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 9:
						v571 = v551
						v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+9)))
						v576 = v572<<(uint(int32(16))%32) + v571
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					case 10:
						v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+10)))
						v571 = v567<<(uint(int32(24))%32) + v551
						v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+9)))
						v576 = v572<<(uint(int32(16))%32) + v571
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+8)))
						v581 = v577<<(uint(int32(8))%32) + v576
						v582 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
						v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
						v676 = v582 + v546
						v677 = v584 + v547
						v678 = v581
					default:
						v676 = v546
						v677 = v547
						v678 = v551
					}
				}
				v681 = int32(14)
				v683 = v677 ^ v678 - base.I32_rotl(v677, v681)
				v687 = v683 ^ v676 - base.I32_rotl(v683, int32(11))
				v691 = v687 ^ v677 - base.I32_rotl(v687, int32(25))
				v695 = v691 ^ v683 - base.I32_rotl(v691, int32(16))
				v699 = v695 ^ v687 - base.I32_rotl(v695, int32(4))
				v703 = v699 ^ v691 - base.I32_rotl(v699, v681)
				*(*int64)(unsafe.Add(mBase, uint32(v382))) = base.I64_extend_i32_u(v703)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v703^v695-base.I32_rotl(v703, int32(24)))
				v715 = int32(8)
			} else {
				v715 = v389
			}
			v717 = int32(1024) - v715
			if base.Ui32(v391) < base.Ui32(v717) {
				v719 = v391
			} else {
				v719 = v717
			}
			if v719 != 0 {
				base.MemoryCopy(m, v715+v382, v388, v719)
			} else {
			}
			v723 = v715 + v719
			v724 = v391 - v719
			if v724 != 0 {
				v388 = v388 + v719
				v389 = v723
				v391 = v724
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v723
		return
	} else {
		v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
		*(*uint16)(unsafe.Add(mBase, uint32(v376+v382))) = uint16(v727)
		v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729 + int32(2)
		return
	}
}
func F_ApplySetting(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v14 = v11 + int32(16)
	F_ScanKeyInit(m, v14, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v11+int32(72), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l2))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = F_systable_beginscan(m, l3, int32(2965), int32(1), l0, int32(2), v14)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = F_systable_getnext(m, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = v34
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_systable_endscan(m, v32)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v47 = F_heap_getattr_4(m, v37, v44, v11+int32(15))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v49 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = F_pg_detoast_datum(m, base.I32_wrap_i64(v47))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v59 = F_systable_getnext(m, v32)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	F_ProcessGUCArray(m, v53, int32(5), l4, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v59 != 0 {
		v37 = v59
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	m.G0 = v11 + int32(128)
	return
}
func F_AutoVacLauncherMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int64
	_ = v200
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int64
	_ = v253
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int64
	_ = v856
	var v857 int64
	_ = v857
	var v867 int32
	_ = v867
	var v870 int64
	_ = v870
	var v877 int64
	_ = v877
	var v881 int64
	_ = v881
	var v882 int64
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int64
	_ = v902
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int64
	_ = v925
	var v926 int64
	_ = v926
	var v936 int32
	_ = v936
	var v939 int64
	_ = v939
	var v946 int64
	_ = v946
	var v950 int64
	_ = v950
	var v951 int64
	_ = v951
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int64
	_ = v974
	var v978 int64
	_ = v978
	var v980 int32
	_ = v980
	var v981 int64
	_ = v981
	var v984 int64
	_ = v984
	var v989 int32
	_ = v989
	var v990 int64
	_ = v990
	var v992 int32
	_ = v992
	var v993 int64
	_ = v993
	var v996 int64
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int64
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1172 int32
	_ = v1172
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int64
	_ = v1212
	var v1213 int64
	_ = v1213
	var v1221 int64
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int64
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1275 int64
	_ = v1275
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1401 int32
	_ = v1401
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1419 int64
	_ = v1419
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1441 int32
	_ = v1441
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1488 int32
	_ = v1488
	var v1501 int32
	_ = v1501
	var v1512 int32
	_ = v1512
	var v1525 int32
	_ = v1525
	var v1535 int32
	_ = v1535
	var v1536 int64
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	v15 = int32(-1)
	v17 = int32(0)
	v18 = v12
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v15 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1535 = int32(m.ExcTag)
	v1536 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1535 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[0]))
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v517 = v17
	goto L9
L9:
	;
	if v517 != 0 {
		goto L121
	} else {
		goto L122
	}
L10:
	;
	F_MemoryContextDelete(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[0])) = int32(0)
	goto L12
L14:
	;
	v41 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L18
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[1]))
	v38 = F_GetBackendTypeDesc(m, v37)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v41 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_errmsg_internal(m, int32(_a_F_AutoVacLauncherMain_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[2]))
	if v53 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(429), int32(_a_F_AutoVacLauncherMain_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_pg_usleep(m, v53*int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	v61 = m.G0
	v63 = v61 - int32(32)
	m.G0 = v63
	v66 = int32(967)
	switch v66 {
	case 0, 2:
		goto L28
	default:
		goto L29
	}
L27:
	;
	v103 = m.G0
	v105 = v103 - int32(32)
	m.G0 = v105
	v108 = int32(968)
	switch v108 {
	case 0, 2:
		goto L38
	default:
		goto L39
	}
L28:
	;
	F_sigemptyset(m, v63+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(268435456)
	switch v66 {
	case 0:
		goto L33
	default:
		goto L31
	case 2:
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[3])) = int32(965)
	goto L28
L30:
	;
	goto L35
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = int32(0)
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = int32(-2)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v95 = F___sigaction(m, int32(1), v63+int32(12), int32(0))
	mBase = m.M
	m.G0 = v63 + int32(32)
	goto L27
L37:
	;
	v145 = m.G0
	v147 = v145 - int32(32)
	m.G0 = v147
	v150 = int32(969)
	switch v150 {
	case 0, 2:
		goto L48
	default:
		goto L49
	}
L38:
	;
	F_sigemptyset(m, v105+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = int32(268435456)
	switch v108 {
	case 0:
		goto L43
	default:
		goto L41
	case 2:
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[4])) = int32(966)
	goto L38
L40:
	;
	goto L45
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(0)
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(-2)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v137 = F___sigaction(m, int32(2), v105+int32(12), int32(0))
	mBase = m.M
	m.G0 = v105 + int32(32)
	goto L37
L47:
	;
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[5])) = v183
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[6])) = v183
	v193 = v183
	goto L58
L48:
	;
	F_sigemptyset(m, v147+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(268435456)
	switch v150 {
	case 0:
		goto L53
	default:
		goto L51
	case 2:
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[7])) = int32(967)
	goto L48
L50:
	;
	goto L55
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = int32(0)
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = int32(-2)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v179 = F___sigaction(m, int32(15), v147+int32(12), int32(0))
	mBase = m.M
	m.G0 = v147 + int32(32)
	goto L47
L57:
	;
	v273 = int32(0)
	v275 = m.G0
	v277 = v275 - int32(32)
	m.G0 = v277
	switch v273 {
	case 0, 2:
		goto L64
	default:
		goto L65
	}
L58:
	;
	v195 = int32(40)
	v196 = v193 * v195
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[8]))) = uint8(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[9]))) = v193
	v200 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[10]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[11]))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[12]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[13]))) = v197
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_AutoVacLauncherMain[14]))) = uint8(v197)
	v211 = v193 | int32(1)
	v213 = v211 * v195
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[8]))) = uint8(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[9]))) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[10]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[11]))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[12]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[13]))) = v197
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+uint32(_c_F_AutoVacLauncherMain[14]))) = uint8(v197)
	v228 = v193 | int32(2)
	v230 = v228 * v195
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[8]))) = uint8(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[9]))) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[10]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[11]))) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[12]))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[13]))) = v197
	*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_c_F_AutoVacLauncherMain[14]))) = uint8(v197)
	if v193 != int32(20) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[15])) = uint8(v266)
	F_pqsignal_be(m, int32(14), int32(1992))
	mBase = m.M
	goto L57
L60:
	;
	v247 = v193 | int32(3)
	v249 = v247 * int32(40)
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[8]))) = uint8(v250)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[9]))) = v247
	v253 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[10]))) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[11]))) = v250
	*(*int64)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[12]))) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[13]))) = v250
	*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_AutoVacLauncherMain[14]))) = uint8(v250)
	v193 = v193 + int32(4)
	goto L58
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v317 = m.G0
	v319 = v317 - int32(32)
	m.G0 = v319
	v322 = int32(970)
	switch v322 {
	case 0, 2:
		goto L74
	default:
		goto L75
	}
L64:
	;
	F_sigemptyset(m, v277+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(268435456)
	switch v273 {
	case 0:
		goto L69
	default:
		goto L67
	case 2:
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[16])) = int32(-2)
	goto L64
L66:
	;
	goto L71
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = int32(0)
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+12)) = int32(-2)
	goto L66
L71:
	;
	goto L72
L72:
	;
	v309 = F___sigaction(m, int32(13), v277+int32(12), int32(0))
	mBase = m.M
	m.G0 = v277 + int32(32)
	goto L63
L73:
	;
	v359 = m.G0
	v361 = v359 - int32(32)
	m.G0 = v361
	v364 = int32(971)
	switch v364 {
	case 0, 2:
		goto L84
	default:
		goto L85
	}
L74:
	;
	F_sigemptyset(m, v319+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v319)+24)) = int32(268435456)
	switch v322 {
	case 0:
		goto L79
	default:
		goto L77
	case 2:
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[17])) = int32(968)
	goto L74
L76:
	;
	goto L81
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+12)) = int32(0)
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+12)) = int32(-2)
	goto L76
L81:
	;
	goto L82
L82:
	;
	v351 = F___sigaction(m, int32(10), v319+int32(12), int32(0))
	mBase = m.M
	m.G0 = v319 + int32(32)
	goto L73
L83:
	;
	v401 = m.G0
	v403 = v401 - int32(32)
	m.G0 = v403
	v406 = int32(972)
	switch v406 {
	case 0, 2:
		goto L94
	default:
		goto L95
	}
L84:
	;
	F_sigemptyset(m, v361+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v361)+24)) = int32(268435456)
	switch v364 {
	case 0:
		goto L89
	default:
		goto L87
	case 2:
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[18])) = int32(969)
	goto L84
L86:
	;
	goto L91
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = int32(0)
	goto L86
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = int32(-2)
	goto L86
L91:
	;
	goto L92
L92:
	;
	v393 = F___sigaction(m, int32(12), v361+int32(12), int32(0))
	mBase = m.M
	m.G0 = v361 + int32(32)
	goto L83
L93:
	;
	v443 = m.G0
	v445 = v443 - int32(32)
	m.G0 = v445
	v447 = int32(2)
	switch v447 {
	case 0, 2:
		goto L104
	default:
		goto L105
	}
L94:
	;
	F_sigemptyset(m, v403+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v403)+24)) = int32(268435456)
	switch v406 {
	case 0:
		goto L99
	default:
		goto L97
	case 2:
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[19])) = int32(970)
	goto L94
L96:
	;
	goto L101
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v403)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+12)) = int32(0)
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+12)) = int32(-2)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v435 = F___sigaction(m, int32(8), v403+int32(12), int32(0))
	mBase = m.M
	m.G0 = v403 + int32(32)
	goto L93
L103:
	;
	F_InitProcess(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L6
	} else {
		goto L113
	}
L104:
	;
	F_sigemptyset(m, v445+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v445)+24)) = int32(268435456)
	switch v447 {
	case 0:
		goto L109
	default:
		goto L107
	case 2:
		goto L108
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[20])) = int32(0)
	goto L104
L106:
	;
	goto L110
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = int32(_a_F_AutoVacLauncherMain_4)
	v470 = int32(268435461)
	goto L106
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = int32(0)
	v470 = int32(268435457)
	goto L106
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = int32(-2)
	v470 = int32(268435457)
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+24)) = v470
	goto L112
L112:
	;
	v477 = F___sigaction(m, int32(17), v445+int32(12), int32(0))
	mBase = m.M
	m.G0 = v445 + int32(32)
	goto L103
L113:
	;
	F_BaseInit(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v485 = int32(0)
	F_InitPostgres(m, v485, v485, v485, v485, v485, v485)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[21])) = int32(2)
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[22]))
	v502 = F_AllocSetContextCreateInternal(m, v497, int32(_a_F_AutoVacLauncherMain_5), int32(0), int32(_a_F_AutoVacLauncherMain_6), int32(_a_F_AutoVacLauncherMain_7))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[23])) = v502
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[24])) = v502
	goto L117
L117:
	;
	v510 = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v510)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v510))) = v18 + int32(28)
	goto L120
L118:
	;
	v517 = int32(0)
	goto L9
L120:
	;
	goto L118
L121:
	;
	v518 = int32(_a_F_AutoVacLauncherMain_8)
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[25])) = v520 + int32(1)
	v525 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[26])) = v525
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[5])) = v525
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[6])) = v525
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[8])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[14])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[27])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[28])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[29])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[30])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[31])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[32])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[33])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[34])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[35])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[36])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[37])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[38])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[39])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[40])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[41])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[42])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[43])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[44])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[45])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[46])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[47])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[48])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[49])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[50])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[51])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[52])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[53])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[54])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[55])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[56])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[57])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[58])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[59])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[60])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[61])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[62])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[63])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[64])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[65])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[66])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[67])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[68])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[69])) = uint8(v525)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[70])) = uint8(v525)
	goto L124
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[71])) = v18 + int32(32)
	F_pgmem_sigprocmask(m, int32(_a_F_AutoVacLauncherMain_9), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L143
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[72])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_LWLockReleaseAll(m)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[73]))
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = int32(0)
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[74]))
	if v689 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L6
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L6
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[23])) = v703
	F_FlushErrorState(m)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[24]))
	F_MemoryContextReset(m, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v711 = int32(_a_F_AutoVacLauncherMain_8)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[25])) = v713 - int32(1)
	v718 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75])) = v718
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76])) = v718
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[77])) = int32(0)
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78]))
	if v727 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	F_pg_usleep(m, int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	goto L123
L142:
	;
	goto L3
L143:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_11), int32(_a_F_AutoVacLauncherMain_12), int32(5), int32(10))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_13), int32(_a_F_AutoVacLauncherMain_14), int32(5), int32(10))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_15), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_17), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_18), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_19), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_20), int32(_a_F_AutoVacLauncherMain_21), int32(5), int32(10))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_22), int32(_a_F_AutoVacLauncherMain_23), int32(5), int32(10))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[79])))
	if v789 == int32(1) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[81]))
	*(*int32)(unsafe.Add(mBase, uint32(v806)+8)) = v808
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L162
	}
L153:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[82])))
	if v793&int32(1) != 0 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78]))
	if v797 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	goto L155
L157:
	;
	v800 = F_do_start_worker(m)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L6
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L6
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	goto L3
L162:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78]))
	if v814 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	goto L166
L164:
	;
	goto L165
L165:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L6
	} else {
		goto L344
	}
L166:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+20))
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83]))
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84]))
	v833 = v830 - v832
	v834 = int32(0)
	if v834 < v833 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L165
L168:
	;
	v902 = base.I64_extend_i32_s(v900)
	if v900 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L169:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85]))
	v900 = v897
	v901 = int32(0)
	goto L168
L170:
	;
	v837 = v833
	goto L172
L171:
	;
	v837 = v834
	goto L172
L172:
	;
	v838 = base.B2i32(v837 < v828)
	if v838 == int32(0) {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v842 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if base.B2i32(v842 == int32(0))|base.B2i32(v842 == int32(_a_F_AutoVacLauncherMain_10)) != 0 {
		goto L169
	} else {
		goto L174
	}
L174:
	;
	v851 = m.G0
	v852 = int32(16)
	v853 = v851 - v852
	m.G0 = v853
	F_gettimeofday(m, v853)
	mBase = m.M
	v856 = *(*int64)(unsafe.Add(mBase, uint32(v853)))
	v857 = int64(*(*int32)(unsafe.Add(mBase, uint32(v853)+8)))
	m.G0 = v853 + v852
	goto L175
L175:
	;
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76]))
	v870 = *(*int64)(unsafe.Add(mBase, uint32(v867-int32(12))))
	v877 = v870 - (v857 + v856*int64(1000000) - int64(946684800000000))
	if v877 <= int64(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v18)+196))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	v900 = v893
	v901 = v894
	goto L168
L177:
	;
	v889 = int32(0)
	v890 = int32(0)
	goto L179
L178:
	;
	v881 = int64(1000000)
	v882 = base.I64_div_u_s(v877, v881)
	v889 = base.I32_wrap_i64(v882)
	v890 = base.I32_wrap_i64(v877 - v882*v881)
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(196)))) = v889
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(192)))) = v890
	goto L176
L180:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[86]))
	v1005 = int32(1000)
	v1008 = base.I32_div_s(v997, v1005)
	v1011 = F_WaitLatch(m, v1002, int32(41), base.I32_wrap_i64(v1000)*v1005+v1008, int32(83886081))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L6
	} else {
		goto L219
	}
L181:
	;
	v993 = int64(300)
	if base.Ui64(v993) <= base.Ui64(v902) {
		goto L216
	} else {
		goto L217
	}
L182:
	;
	v989 = base.B2i32(v901 < int32(_a_F_AutoVacLauncherMain_24))
	if v901 < int32(_a_F_AutoVacLauncherMain_24) {
		goto L210
	} else {
		goto L211
	}
L183:
	;
	if v901 != 0 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	if int32(0) < v900 {
		goto L181
	} else {
		goto L209
	}
L186:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	if v838 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v974 = base.I64_extend_i32_s(v973)
	if v973 <= int32(0) {
		goto L197
	} else {
		goto L198
	}
L189:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85]))
	v971 = int32(0)
	v972 = int32(1)
	v973 = v970
	goto L188
L190:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if base.B2i32(v911 == int32(0))|base.B2i32(v911 == int32(_a_F_AutoVacLauncherMain_10)) != 0 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v920 = m.G0
	v921 = int32(16)
	v922 = v920 - v921
	m.G0 = v922
	F_gettimeofday(m, v922)
	mBase = m.M
	v925 = *(*int64)(unsafe.Add(mBase, uint32(v922)))
	v926 = int64(*(*int32)(unsafe.Add(mBase, uint32(v922)+8)))
	m.G0 = v922 + v921
	goto L192
L192:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76]))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v936-int32(12))))
	v946 = v939 - (v926 + v925*int64(1000000) - int64(946684800000000))
	if v946 <= int64(0) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	v971 = v962
	v972 = base.B2i32(v962 < int32(_a_F_AutoVacLauncherMain_24))
	v973 = v965
	goto L188
L194:
	;
	v958 = int32(0)
	v959 = int32(0)
	goto L196
L195:
	;
	v950 = int64(1000000)
	v951 = base.I64_div_u_s(v946, v950)
	v958 = base.I32_wrap_i64(v951)
	v959 = base.I32_wrap_i64(v946 - v951*v950)
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(204)))) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(200)))) = v959
	goto L193
L197:
	;
	if v972 != 0 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v981 = int64(300)
	if base.Ui64(v981) <= base.Ui64(v974) {
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v978 = int64(0)
	goto L202
L201:
	;
	v978 = v974
	goto L202
L202:
	;
	if v972 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v980 = int32(_a_F_AutoVacLauncherMain_25)
	goto L205
L204:
	;
	v980 = v971
	goto L205
L205:
	;
	v997 = v980
	v1000 = v978
	goto L180
L206:
	;
	v984 = v981
	goto L208
L207:
	;
	v984 = v974
	goto L208
L208:
	;
	v997 = v971
	v1000 = v984
	goto L180
L209:
	;
	goto L182
L210:
	;
	v990 = int64(0)
	goto L212
L211:
	;
	v990 = v902
	goto L212
L212:
	;
	if v901 < int32(_a_F_AutoVacLauncherMain_24) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v992 = int32(_a_F_AutoVacLauncherMain_25)
	goto L215
L214:
	;
	v992 = v901
	goto L215
L215:
	;
	v997 = v992
	v1000 = v990
	goto L180
L216:
	;
	v996 = v993
	goto L218
L217:
	;
	v996 = v902
	goto L218
L218:
	;
	v997 = v901
	v1000 = v996
	goto L180
L219:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[86]))
	v1015 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1014))) = v1015
	v1020 = base.AtomicRmwOr32(m, v1015, int32(_a_F_AutoVacLauncherMain_26), v1015)
	goto L220
L220:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78]))
	if v1022 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L6
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[87]))
	if v1026 != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L3
L225:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[87])) = int32(0)
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84]))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L6
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[88]))
	if v1089 != 0 {
		goto L245
	} else {
		goto L246
	}
L228:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[79])))
	if v1036 == int32(1) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84]))
	if v1031 == v1046 {
		goto L235
	} else {
		goto L236
	}
L230:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[82])))
	if v1040&int32(1) != 0 {
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L6
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	goto L3
L235:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L6
	} else {
		goto L244
	}
L236:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83]))
	if v1046 <= v1049 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1053 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L6
	} else {
		goto L238
	}
L238:
	;
	if v1053 == int32(0) {
		goto L235
	} else {
		goto L239
	}
L239:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84]))
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1061
	F_errmsg(m, int32(_a_F_AutoVacLauncherMain_27), v18+int32(16))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L6
	} else {
		goto L241
	}
L241:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1072
	v1075 = F_errdetail(m, int32(_a_F_AutoVacLauncherMain_28), v18)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L6
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(3640), int32(_a_F_AutoVacLauncherMain_29))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L6
	} else {
		goto L243
	}
L243:
	;
	goto L235
L244:
	;
	goto L227
L245:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L6
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[89]))
	if v1093 != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	goto L247
L249:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L6
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L6
	} else {
		goto L253
	}
L252:
	;
	goto L251
L253:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[90]))
	if v1099 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78]))
	if v1512 == int32(0) {
		goto L166
	} else {
		goto L343
	}
L255:
	;
	v1207 = m.G0
	v1208 = int32(16)
	v1209 = v1207 - v1208
	m.G0 = v1209
	F_gettimeofday(m, v1209)
	mBase = m.M
	v1212 = *(*int64)(unsafe.Add(mBase, uint32(v1209)))
	v1213 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1209)+8)))
	m.G0 = v1209 + v1208
	v1221 = v1213 + v1212*int64(1000000) - int64(946684800000000)
	goto L279
L256:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[90])) = int32(0)
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1107 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	v1113 = F_LWLockAcquire(m, v1109+int32(2816), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L6
	} else {
		goto L260
	}
L258:
	;
	v1163 = v1106
	goto L259
L259:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1163)))
	if v1172 == int32(0) {
		goto L255
	} else {
		goto L274
	}
L260:
	;
	v1115 = int32(0)
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+4)) = v1115
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+uint32(_c_F_AutoVacLauncherMain[92])))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+28))
	if v1121 == v1115 {
		v1146 = v1115
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if v1146 != v1120 {
		goto L270
	} else {
		goto L271
	}
L262:
	;
	v1125 = v1117 + int32(24)
	if v1121 == v1125 {
		v1146 = v1115
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1127 = v1121
	v1129 = v1115
	goto L264
L264:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+16))
	if v1136 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1146 = v1141
	goto L261
L266:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+32))
	v1141 = v1129 + base.B2i32(v1137 != int32(0))
	goto L268
L267:
	;
	v1141 = v1129
	goto L268
L268:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	if v1142 != v1125 {
		v1127 = v1142
		v1129 = v1141
		goto L264
	} else {
		goto L269
	}
L269:
	;
	goto L265
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+uint32(_c_F_AutoVacLauncherMain[92]))) = v1146
	goto L272
L271:
	;
	goto L272
L272:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	F_LWLockRelease(m, v1156+int32(2816))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L6
	} else {
		goto L273
	}
L273:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v1163 = v1162
	goto L259
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1163))) = int32(0)
	F_pg_usleep(m, int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[93])))
	if v1181 == int32(1) {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	goto L254
L276:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[94]))
	*(*int32)(unsafe.Add(mBase, uint32(v1185+int32(20)))) = int32(1)
	v1192 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[95]))
	v1194 = F_pgmem_kill(m, v1192, int32(10))
	mBase = m.M
	goto L278
L277:
	;
	goto L278
L278:
	;
	goto L275
L279:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	v1227 = F_LWLockAcquire(m, v1223+int32(2816), int32(1))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L6
	} else {
		goto L280
	}
L280:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84]))
	v1232 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83]))
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+20))
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+32))
	if v1236 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	F_LWLockRelease(m, v1322+int32(2816))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L6
	} else {
		goto L298
	}
L282:
	;
	v1239 = *(*int64)(unsafe.Add(mBase, uint32(v1236)+24))
	v1240 = int32(60)
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85]))
	if v1240 <= v1242 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1245 = v1240
	goto L285
L284:
	;
	v1245 = v1242
	goto L285
L285:
	;
	goto L286
L286:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	F_LWLockRelease(m, v1254+int32(2816))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L6
	} else {
		goto L287
	}
L287:
	;
	if base.B2i32(base.I64_extend_i32_s(v1245*int32(1000))*int64(1000) <= v1221-v1239) == int32(0) {
		goto L254
	} else {
		goto L288
	}
L288:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91]))
	v1266 = F_LWLockAcquire(m, v1262+int32(2816), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L6
	} else {
		goto L289
	}
L289:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+32))
	if v1270 == int32(0) {
		goto L281
	} else {
		goto L290
	}
L290:
	;
	v1273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1270)+36)) = uint8(v1273)
	v1275 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1270)+8)) = v1275
	*(*int64)(unsafe.Add(mBase, uint32(v1270)+24)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+16)) = v1273
	v1282 = v1269 + int32(12)
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+16))
	if v1283 == v1273 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+12)) = v1282
	v1289 = v1282
	goto L293
L292:
	;
	v1289 = v1283
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1270))) = v1282
	*(*int32)(unsafe.Add(mBase, uint32(v1270)+4)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v1289))) = v1270
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+16)) = v1270
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+20)) = v1294 + int32(1)
	v1299 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80]))
	v1300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1299)+32)) = v1300
	v1304 = F_errstart(m, int32(19), v1300)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	if v1304 == int32(0) {
		goto L281
	} else {
		goto L295
	}
L295:
	;
	F_errmsg(m, int32(_a_F_AutoVacLauncherMain_30), int32(0))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L6
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(737), int32(_a_F_AutoVacLauncherMain_2))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L6
	} else {
		goto L297
	}
L297:
	;
	goto L281
L298:
	;
	v1327 = v1232 - v1230
	v1328 = int32(0)
	if v1328 < v1327 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1331 = v1327
	goto L301
L300:
	;
	v1331 = v1328
	goto L301
L301:
	;
	if v1235 <= v1331 {
		goto L254
	} else {
		goto L302
	}
L302:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if v1334 != int32(_a_F_AutoVacLauncherMain_10) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1338 = v1334
	goto L305
L304:
	;
	v1338 = int32(0)
	goto L305
L305:
	;
	if v1338 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1341 = F_do_start_worker(m)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L6
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76]))
	v1419 = *(*int64)(unsafe.Add(mBase, uint32(v1416-int32(12))))
	goto L325
L309:
	;
	if v1341 == int32(0) {
		goto L254
	} else {
		goto L310
	}
L310:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	v1347 = int32(0)
	if base.B2i32(v1346 == v1347)|base.B2i32(v1346 == int32(_a_F_AutoVacLauncherMain_10)) == v1347 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1354 = v1346
	goto L314
L312:
	;
	goto L313
L313:
	;
	F_rebuild_database_list(m, v1341)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L6
	} else {
		goto L324
	}
L314:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1354-int32(20))))
	if v1341 == v1365 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	goto L313
L316:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85]))
	*(*int64)(unsafe.Add(mBase, uint32(v1354-int32(12)))) = base.I64_extend_i32_s(v1370*int32(1000))*int64(1000) + v1221
	v1379 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if v1379 == v1354 {
		goto L254
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+4))
	if v1401 != int32(_a_F_AutoVacLauncherMain_10) {
		v1354 = v1401
		goto L314
	} else {
		goto L323
	}
L319:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1381)+4)) = v1382
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1354)))
	*(*int32)(unsafe.Add(mBase, uint32(v1382))) = v1384
	v1387 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if v1387 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1390 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76])) = v1390
	v1394 = v1390
	goto L322
L321:
	;
	v1394 = v1387
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1354))) = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, uint32(v1354)+4)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v1394))) = v1354
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75])) = v1354
	goto L254
L323:
	;
	goto L315
L324:
	;
	goto L254
L325:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(0))*int64(1000) <= v1221-v1419) == int32(0) {
		goto L254
	} else {
		goto L326
	}
L326:
	;
	v1428 = F_do_start_worker(m)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L6
	} else {
		goto L327
	}
L327:
	;
	if v1428 == int32(0) {
		goto L254
	} else {
		goto L328
	}
L328:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	v1434 = int32(0)
	if base.B2i32(v1433 == v1434)|base.B2i32(v1433 == int32(_a_F_AutoVacLauncherMain_10)) == v1434 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1441 = v1433
	goto L332
L330:
	;
	goto L331
L331:
	;
	F_rebuild_database_list(m, v1428)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L6
	} else {
		goto L342
	}
L332:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1441-int32(20))))
	if v1428 == v1452 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L331
L334:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85]))
	*(*int64)(unsafe.Add(mBase, uint32(v1441-int32(12)))) = base.I64_extend_i32_s(v1457*int32(1000))*int64(1000) + v1221
	v1466 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if v1466 == v1441 {
		goto L254
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	if v1488 != int32(_a_F_AutoVacLauncherMain_10) {
		v1441 = v1488
		goto L332
	} else {
		goto L341
	}
L337:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1441)))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1468)+4)) = v1469
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1441)))
	*(*int32)(unsafe.Add(mBase, uint32(v1469))) = v1471
	v1474 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75]))
	if v1474 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1477 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76])) = v1477
	v1481 = v1477
	goto L340
L339:
	;
	v1481 = v1474
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1441))) = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, uint32(v1441)+4)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(v1481))) = v1441
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75])) = v1441
	goto L254
L341:
	;
	goto L333
L342:
	;
	goto L254
L343:
	;
	goto L167
L344:
	;
	goto L5
L345:
	;
	v1540 = int32(v1536)
	m.G0 = v18
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+4))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1540)))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1543)))
	if v18+int32(28) == v1546 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	m.ExcPending = 1
	goto L354
L347:
	;
	if v1550 == int32(0) {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+4))
	v1550 = v1548
	goto L350
L349:
	;
	v1550 = int32(0)
	goto L350
L350:
	;
	goto L347
L351:
	;
	F___wasm_longjmp(m, v1543, v1542)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	goto L353
L353:
	;
	v15 = v1550
	v17 = v1542
	goto L1
L354:
	;
	return
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AuxiliaryProcKill(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v7 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	if v6 == v7 {
		F_LWLockReleaseAll(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_ConditionVariableCancelSleep(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_SwitchBackToLocalLatch(m)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[1])) = int32(_a_F_AuxiliaryProcKill_0)
					*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[2])) = int32(-1)
					v21 = int32(_a_F_AuxiliaryProcKill_1)
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[0]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[0])) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v22+int32(316))+12)) = v24
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
					v34 = base.AtomicRmwXchg32(m, v31, int32(20), int32(1))
					if v34 != 0 {
						F_s_lock(m, v31+int32(20), int32(_a_F_AuxiliaryProcKill_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
							v53 = base.I32_div_s(v48+v46*int32(15), int32(16))
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v55)+72)) = v53
							v57 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v55)+20)), uint32(v57))
							return
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = int64(4294967295)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
						v53 = base.I32_div_s(v48+v46*int32(15), int32(16))
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v55)+72)) = v53
						v57 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v55)+20)), uint32(v57))
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_AuxiliaryProcKill_3), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_AuxiliaryProcKill_4), int32(1083), int32(_a_F_AuxiliaryProcKill_5))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F___addtf3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v140 int64
	_ = v140
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v186 int64
	_ = v186
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v212 int64
	_ = v212
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v241 int64
	_ = v241
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v255 int32
	_ = v255
	var v270 int64
	_ = v270
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v279 int64
	_ = v279
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v297 int64
	_ = v297
	var v304 int64
	_ = v304
	var v308 int64
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v334 int64
	_ = v334
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v359 int64
	_ = v359
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v374 int32
	_ = v374
	var v377 int64
	_ = v377
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v400 int64
	_ = v400
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v414 int32
	_ = v414
	var v429 int64
	_ = v429
	var v433 int64
	_ = v433
	var v434 int64
	_ = v434
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v446 int64
	_ = v446
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int32
	_ = v449
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
	var v463 int64
	_ = v463
	var v466 int32
	_ = v466
	var v472 int64
	_ = v472
	var v475 int64
	_ = v475
	var v478 int64
	_ = v478
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v488 int64
	_ = v488
	var v489 int64
	_ = v489
	v6 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = int64(9223372036854775807)
	v19 = l4 & v18
	v21 = base.B2i32(l1 == v6)
	v23 = l2 & v18
	if v23 == v6 {
		v30 = v21
	} else {
		v30 = base.B2i32(base.Ui64(v23-int64(9223090561878065152)) < base.Ui64(int64(-9223090561878065152)))
	}
	if v30 == int32(0) {
		v36 = v19 - int64(9223090561878065152)
		v37 = int64(-9223090561878065152)
		if v36 == v37 {
			v41 = base.B2i32(l3 != int64(0))
		} else {
			v41 = base.B2i32(base.Ui64(v37) < base.Ui64(v36))
		}
		if v41 != 0 {
			if v19 == v23 {
				v99 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
			} else {
				v99 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
			}
			if v99 != 0 {
				v100 = l3
			} else {
				v100 = l1
			}
			if v99 != 0 {
				v101 = l4
			} else {
				v101 = l2
			}
			v103 = v101 & int64(281474976710655)
			if v99 != 0 {
				v104 = l2
			} else {
				v104 = l4
			}
			v105 = int64(48)
			v108 = int32(_a_F___addtf3_0)
			v109 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v105)%64))) & v108
			v114 = base.I32_wrap_i64(int64(base.Ui64(v101)>>(uint(v105)%64))) & v108
			if v114 == int32(0) {
				v118 = v16 + int32(96)
				v120 = base.B2i32(v103 == int64(0))
				if v103 == int64(0) {
					v121 = v100
				} else {
					v121 = v103
				}
				if v103 == int64(0) {
					v125 = int64(64)
				} else {
					v125 = int64(0)
				}
				v127 = base.I32_wrap_i64(base.I64_clz(v121) + v125)
				v129 = v127 - int32(15)
				if v129&int32(64) != 0 {
					v148 = int64(0)
					v149 = v100 << (uint(base.I64_extend_i32_u(v129+int32(-64))) % 64)
				} else {
					if v129 == int32(0) {
						v148 = v100
						v149 = v103
					} else {
						v140 = base.I64_extend_i32_u(v129)
						v148 = v100 << (uint(v140) % 64)
						v149 = v103<<(uint(v140)%64) | int64(base.Ui64(v100)>>(uint(base.I64_extend_i32_u(int32(64)-v129))%64))
					}
				}
				*(*int64)(unsafe.Add(mBase, uint32(v118))) = v148
				*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v149
				v153 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
				v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
				v157 = v154
				v158 = v153
				v159 = int32(16) - v127
			} else {
				v157 = v103
				v158 = v100
				v159 = v114
			}
			if v99 != 0 {
				v160 = l1
			} else {
				v160 = l3
			}
			v162 = v104 & int64(281474976710655)
			if v109 != 0 {
				v203 = v160
				v204 = v109
				v205 = v162
			} else {
				v164 = v16 + int32(80)
				v166 = base.B2i32(v162 == int64(0))
				if v162 == int64(0) {
					v167 = v160
				} else {
					v167 = v162
				}
				if v162 == int64(0) {
					v171 = int64(64)
				} else {
					v171 = int64(0)
				}
				v173 = base.I32_wrap_i64(base.I64_clz(v167) + v171)
				v175 = v173 - int32(15)
				if v175&int32(64) != 0 {
					v194 = int64(0)
					v195 = v160 << (uint(base.I64_extend_i32_u(v175+int32(-64))) % 64)
				} else {
					if v175 == int32(0) {
						v194 = v160
						v195 = v162
					} else {
						v186 = base.I64_extend_i32_u(v175)
						v194 = v160 << (uint(v186) % 64)
						v195 = v162<<(uint(v186)%64) | int64(base.Ui64(v160)>>(uint(base.I64_extend_i32_u(int32(64)-v175))%64))
					}
				}
				*(*int64)(unsafe.Add(mBase, uint32(v164))) = v194
				*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = v195
				v201 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
				v202 = *(*int64)(unsafe.Add(mBase, uint32(v16)+88))
				v203 = v201
				v204 = int32(16) - v173
				v205 = v202
			}
			v206 = int64(3)
			v208 = int64(61)
			v212 = v205<<(uint(v206)%64) | int64(base.Ui64(v203)>>(uint(v208)%64)) | int64(2251799813685248)
			v220 = v203 << (uint(v206) % 64)
			if v159 == v204 {
				v288 = v212
				v290 = v220
			} else {
				v222 = v159 - v204
				if base.Ui32(int32(127)) < base.Ui32(v222) {
					v288 = int64(0)
					v290 = int64(1)
				} else {
					v228 = v16 - int32(-64)
					v230 = int32(128) - v222
					if v230&int32(64) != 0 {
						v249 = int64(0)
						v250 = v220 << (uint(base.I64_extend_i32_u(v230+int32(-64))) % 64)
					} else {
						if v230 == int32(0) {
							v249 = v220
							v250 = v212
						} else {
							v241 = base.I64_extend_i32_u(v230)
							v249 = v220 << (uint(v241) % 64)
							v250 = v212<<(uint(v241)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v230))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v228))) = v249
					*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = v250
					v255 = v16 + int32(48)
					if v222&int32(64) != 0 {
						v274 = int64(base.Ui64(v212) >> (uint(base.I64_extend_i32_u(v222+int32(-64))) % 64))
						v275 = int64(0)
					} else {
						if v222 == int32(0) {
							v274 = v220
							v275 = v212
						} else {
							v270 = base.I64_extend_i32_u(v222)
							v274 = v212<<(uint(base.I64_extend_i32_u(int32(64)-v222))%64) | int64(base.Ui64(v220)>>(uint(v270)%64))
							v275 = int64(base.Ui64(v212) >> (uint(v270) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v255))) = v274
					*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v275
					v279 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
					v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
					v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
					v288 = v279
					v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
				}
			}
			v292 = v157<<(uint(v206)%64) | int64(base.Ui64(v158)>>(uint(v208)%64)) | int64(2251799813685248)
			v294 = v158 << (uint(int64(3)) % 64)
			if l2^l4 < int64(0) {
				v297 = int64(0)
				if v290^v294|(v288^v292) == v297 {
					v488 = v297
					v489 = v297
				} else {
					v304 = v294 - v290
					v308 = v292 - v288 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v294) < base.Ui64(v290)))
					if base.Ui64(int64(2251799813685247)) < base.Ui64(v308) {
						v371 = v304
						v373 = v308
						v374 = v159
					} else {
						v312 = v16 + int32(32)
						v314 = base.B2i32(v308 == int64(0))
						if v308 == int64(0) {
							v315 = v304
						} else {
							v315 = v308
						}
						if v308 == int64(0) {
							v319 = int64(64)
						} else {
							v319 = int64(0)
						}
						v323 = base.I32_wrap_i64(base.I64_clz(v315)|v319) - int32(12)
						if v323&int32(64) != 0 {
							v342 = int64(0)
							v343 = v304 << (uint(base.I64_extend_i32_u(v323+int32(-64))) % 64)
						} else {
							if v323 == int32(0) {
								v342 = v304
								v343 = v308
							} else {
								v334 = base.I64_extend_i32_u(v323)
								v342 = v304 << (uint(v334) % 64)
								v343 = v308<<(uint(v334)%64) | int64(base.Ui64(v304)>>(uint(base.I64_extend_i32_u(int32(64)-v323))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v312))) = v342
						*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = v343
						v348 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
						v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
						v371 = v349
						v373 = v348
						v374 = v159 - v323
					}
					v377 = v101 & int64(-9223372036854775807-1)
					if int32(_a_F___addtf3_0) <= v374 {
						v488 = int64(0)
						v489 = v377 | int64(9223090561878065152)
					} else {
						v383 = int32(0)
						if v383 < v374 {
							v447 = v371
							v448 = v373
							v449 = v374
						} else {
							v387 = v16 + int32(16)
							v389 = v374 + int32(127)
							if v389&int32(64) != 0 {
								v408 = int64(0)
								v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
							} else {
								if v389 == int32(0) {
									v408 = v371
									v409 = v373
								} else {
									v400 = base.I64_extend_i32_u(v389)
									v408 = v371 << (uint(v400) % 64)
									v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
							*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
							v414 = int32(1) - v374
							if v414&int32(64) != 0 {
								v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
								v434 = int64(0)
							} else {
								if v414 == int32(0) {
									v433 = v371
									v434 = v373
								} else {
									v429 = base.I64_extend_i32_u(v414)
									v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
									v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
							v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
							v448 = v446
							v449 = v383
						}
						v452 = int64(3)
						v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
						v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
						v466 = base.I32_wrap_i64(v447) & int32(7)
						if v466 != int32(4) {
							v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
							v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
							if v466 == int32(0) {
								v488 = v472
								v489 = v475
							} else {
								v484 = v472
								v485 = v475
								v488 = v484
								v489 = v485
							}
						} else {
							v478 = v454 + v454&int64(1)
							v484 = v478
							v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
							v488 = v484
							v489 = v485
						}
					}
				}
			} else {
				v350 = v290 + v294
				v354 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v350) < base.Ui64(v290))) + (v288 + v292)
				if v354&int64(4503599627370496) == int64(0) {
					v371 = v350
					v373 = v354
					v374 = v159
				} else {
					v359 = int64(1)
					v371 = v290&v359 | (v354<<(uint(int64(63))%64) | int64(base.Ui64(v350)>>(uint(v359)%64)))
					v373 = int64(base.Ui64(v354) >> (uint(v359) % 64))
					v374 = v159 + int32(1)
				}
				v377 = v101 & int64(-9223372036854775807-1)
				if int32(_a_F___addtf3_0) <= v374 {
					v488 = int64(0)
					v489 = v377 | int64(9223090561878065152)
				} else {
					v383 = int32(0)
					if v383 < v374 {
						v447 = v371
						v448 = v373
						v449 = v374
					} else {
						v387 = v16 + int32(16)
						v389 = v374 + int32(127)
						if v389&int32(64) != 0 {
							v408 = int64(0)
							v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
						} else {
							if v389 == int32(0) {
								v408 = v371
								v409 = v373
							} else {
								v400 = base.I64_extend_i32_u(v389)
								v408 = v371 << (uint(v400) % 64)
								v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
						*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
						v414 = int32(1) - v374
						if v414&int32(64) != 0 {
							v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
							v434 = int64(0)
						} else {
							if v414 == int32(0) {
								v433 = v371
								v434 = v373
							} else {
								v429 = base.I64_extend_i32_u(v414)
								v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
								v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
						v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
						v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
						v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
						v448 = v446
						v449 = v383
					}
					v452 = int64(3)
					v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
					v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
					v466 = base.I32_wrap_i64(v447) & int32(7)
					if v466 != int32(4) {
						v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
						v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
						if v466 == int32(0) {
							v488 = v472
							v489 = v475
						} else {
							v484 = v472
							v485 = v475
							v488 = v484
							v489 = v485
						}
					} else {
						v478 = v454 + v454&int64(1)
						v484 = v478
						v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
						v488 = v484
						v489 = v485
					}
				}
			}
		} else {
			v43 = int64(9223090561878065152)
			if v23 == v43 {
				v47 = v21
			} else {
				v47 = base.B2i32(base.Ui64(v23) < base.Ui64(v43))
			}
			if v47 == int32(0) {
				v488 = l1
				v489 = l2 | int64(140737488355328)
			} else {
				v54 = int64(9223090561878065152)
				if v19 == v54 {
					v58 = base.B2i32(l3 == int64(0))
				} else {
					v58 = base.B2i32(base.Ui64(v19) < base.Ui64(v54))
				}
				if v58 == int32(0) {
					v488 = l3
					v489 = l4 | int64(140737488355328)
				} else {
					if l1|(v23^int64(9223090561878065152)) == int64(0) {
						v75 = base.B2i32(l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0))
						if l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0) {
							v76 = int64(9223231299366420480)
						} else {
							v76 = l2
						}
						if l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0) {
							v78 = int64(0)
						} else {
							v78 = l1
						}
						v488 = v78
						v489 = v76
					} else {
						if l3|(v19^int64(9223090561878065152)) == int64(0) {
							v488 = l3
							v489 = l4
						} else {
							if l1|v23 == int64(0) {
								if l3|v19 != int64(0) {
									v488 = l3
									v489 = l4
								} else {
									v488 = l1 & l3
									v489 = l2 & l4
								}
							} else {
								if l3|v19 != int64(0) {
									if v19 == v23 {
										v99 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
									} else {
										v99 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
									}
									if v99 != 0 {
										v100 = l3
									} else {
										v100 = l1
									}
									if v99 != 0 {
										v101 = l4
									} else {
										v101 = l2
									}
									v103 = v101 & int64(281474976710655)
									if v99 != 0 {
										v104 = l2
									} else {
										v104 = l4
									}
									v105 = int64(48)
									v108 = int32(_a_F___addtf3_0)
									v109 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v105)%64))) & v108
									v114 = base.I32_wrap_i64(int64(base.Ui64(v101)>>(uint(v105)%64))) & v108
									if v114 == int32(0) {
										v118 = v16 + int32(96)
										v120 = base.B2i32(v103 == int64(0))
										if v103 == int64(0) {
											v121 = v100
										} else {
											v121 = v103
										}
										if v103 == int64(0) {
											v125 = int64(64)
										} else {
											v125 = int64(0)
										}
										v127 = base.I32_wrap_i64(base.I64_clz(v121) + v125)
										v129 = v127 - int32(15)
										if v129&int32(64) != 0 {
											v148 = int64(0)
											v149 = v100 << (uint(base.I64_extend_i32_u(v129+int32(-64))) % 64)
										} else {
											if v129 == int32(0) {
												v148 = v100
												v149 = v103
											} else {
												v140 = base.I64_extend_i32_u(v129)
												v148 = v100 << (uint(v140) % 64)
												v149 = v103<<(uint(v140)%64) | int64(base.Ui64(v100)>>(uint(base.I64_extend_i32_u(int32(64)-v129))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v118))) = v148
										*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v149
										v153 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
										v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
										v157 = v154
										v158 = v153
										v159 = int32(16) - v127
									} else {
										v157 = v103
										v158 = v100
										v159 = v114
									}
									if v99 != 0 {
										v160 = l1
									} else {
										v160 = l3
									}
									v162 = v104 & int64(281474976710655)
									if v109 != 0 {
										v203 = v160
										v204 = v109
										v205 = v162
									} else {
										v164 = v16 + int32(80)
										v166 = base.B2i32(v162 == int64(0))
										if v162 == int64(0) {
											v167 = v160
										} else {
											v167 = v162
										}
										if v162 == int64(0) {
											v171 = int64(64)
										} else {
											v171 = int64(0)
										}
										v173 = base.I32_wrap_i64(base.I64_clz(v167) + v171)
										v175 = v173 - int32(15)
										if v175&int32(64) != 0 {
											v194 = int64(0)
											v195 = v160 << (uint(base.I64_extend_i32_u(v175+int32(-64))) % 64)
										} else {
											if v175 == int32(0) {
												v194 = v160
												v195 = v162
											} else {
												v186 = base.I64_extend_i32_u(v175)
												v194 = v160 << (uint(v186) % 64)
												v195 = v162<<(uint(v186)%64) | int64(base.Ui64(v160)>>(uint(base.I64_extend_i32_u(int32(64)-v175))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v164))) = v194
										*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = v195
										v201 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
										v202 = *(*int64)(unsafe.Add(mBase, uint32(v16)+88))
										v203 = v201
										v204 = int32(16) - v173
										v205 = v202
									}
									v206 = int64(3)
									v208 = int64(61)
									v212 = v205<<(uint(v206)%64) | int64(base.Ui64(v203)>>(uint(v208)%64)) | int64(2251799813685248)
									v220 = v203 << (uint(v206) % 64)
									if v159 == v204 {
										v288 = v212
										v290 = v220
									} else {
										v222 = v159 - v204
										if base.Ui32(int32(127)) < base.Ui32(v222) {
											v288 = int64(0)
											v290 = int64(1)
										} else {
											v228 = v16 - int32(-64)
											v230 = int32(128) - v222
											if v230&int32(64) != 0 {
												v249 = int64(0)
												v250 = v220 << (uint(base.I64_extend_i32_u(v230+int32(-64))) % 64)
											} else {
												if v230 == int32(0) {
													v249 = v220
													v250 = v212
												} else {
													v241 = base.I64_extend_i32_u(v230)
													v249 = v220 << (uint(v241) % 64)
													v250 = v212<<(uint(v241)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v230))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v228))) = v249
											*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = v250
											v255 = v16 + int32(48)
											if v222&int32(64) != 0 {
												v274 = int64(base.Ui64(v212) >> (uint(base.I64_extend_i32_u(v222+int32(-64))) % 64))
												v275 = int64(0)
											} else {
												if v222 == int32(0) {
													v274 = v220
													v275 = v212
												} else {
													v270 = base.I64_extend_i32_u(v222)
													v274 = v212<<(uint(base.I64_extend_i32_u(int32(64)-v222))%64) | int64(base.Ui64(v220)>>(uint(v270)%64))
													v275 = int64(base.Ui64(v212) >> (uint(v270) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v255))) = v274
											*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v275
											v279 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
											v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
											v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
											v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
											v288 = v279
											v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
										}
									}
									v292 = v157<<(uint(v206)%64) | int64(base.Ui64(v158)>>(uint(v208)%64)) | int64(2251799813685248)
									v294 = v158 << (uint(int64(3)) % 64)
									if l2^l4 < int64(0) {
										v297 = int64(0)
										if v290^v294|(v288^v292) == v297 {
											v488 = v297
											v489 = v297
										} else {
											v304 = v294 - v290
											v308 = v292 - v288 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v294) < base.Ui64(v290)))
											if base.Ui64(int64(2251799813685247)) < base.Ui64(v308) {
												v371 = v304
												v373 = v308
												v374 = v159
											} else {
												v312 = v16 + int32(32)
												v314 = base.B2i32(v308 == int64(0))
												if v308 == int64(0) {
													v315 = v304
												} else {
													v315 = v308
												}
												if v308 == int64(0) {
													v319 = int64(64)
												} else {
													v319 = int64(0)
												}
												v323 = base.I32_wrap_i64(base.I64_clz(v315)|v319) - int32(12)
												if v323&int32(64) != 0 {
													v342 = int64(0)
													v343 = v304 << (uint(base.I64_extend_i32_u(v323+int32(-64))) % 64)
												} else {
													if v323 == int32(0) {
														v342 = v304
														v343 = v308
													} else {
														v334 = base.I64_extend_i32_u(v323)
														v342 = v304 << (uint(v334) % 64)
														v343 = v308<<(uint(v334)%64) | int64(base.Ui64(v304)>>(uint(base.I64_extend_i32_u(int32(64)-v323))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v312))) = v342
												*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = v343
												v348 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
												v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
												v371 = v349
												v373 = v348
												v374 = v159 - v323
											}
											v377 = v101 & int64(-9223372036854775807-1)
											if int32(_a_F___addtf3_0) <= v374 {
												v488 = int64(0)
												v489 = v377 | int64(9223090561878065152)
											} else {
												v383 = int32(0)
												if v383 < v374 {
													v447 = v371
													v448 = v373
													v449 = v374
												} else {
													v387 = v16 + int32(16)
													v389 = v374 + int32(127)
													if v389&int32(64) != 0 {
														v408 = int64(0)
														v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
													} else {
														if v389 == int32(0) {
															v408 = v371
															v409 = v373
														} else {
															v400 = base.I64_extend_i32_u(v389)
															v408 = v371 << (uint(v400) % 64)
															v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
														}
													}
													*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
													*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
													v414 = int32(1) - v374
													if v414&int32(64) != 0 {
														v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
														v434 = int64(0)
													} else {
														if v414 == int32(0) {
															v433 = v371
															v434 = v373
														} else {
															v429 = base.I64_extend_i32_u(v414)
															v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
															v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
														}
													}
													*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
													*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
													v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
													v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
													v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
													v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
													v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
													v448 = v446
													v449 = v383
												}
												v452 = int64(3)
												v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
												v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
												v466 = base.I32_wrap_i64(v447) & int32(7)
												if v466 != int32(4) {
													v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
													v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
													if v466 == int32(0) {
														v488 = v472
														v489 = v475
													} else {
														v484 = v472
														v485 = v475
														v488 = v484
														v489 = v485
													}
												} else {
													v478 = v454 + v454&int64(1)
													v484 = v478
													v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
													v488 = v484
													v489 = v485
												}
											}
										}
									} else {
										v350 = v290 + v294
										v354 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v350) < base.Ui64(v290))) + (v288 + v292)
										if v354&int64(4503599627370496) == int64(0) {
											v371 = v350
											v373 = v354
											v374 = v159
										} else {
											v359 = int64(1)
											v371 = v290&v359 | (v354<<(uint(int64(63))%64) | int64(base.Ui64(v350)>>(uint(v359)%64)))
											v373 = int64(base.Ui64(v354) >> (uint(v359) % 64))
											v374 = v159 + int32(1)
										}
										v377 = v101 & int64(-9223372036854775807-1)
										if int32(_a_F___addtf3_0) <= v374 {
											v488 = int64(0)
											v489 = v377 | int64(9223090561878065152)
										} else {
											v383 = int32(0)
											if v383 < v374 {
												v447 = v371
												v448 = v373
												v449 = v374
											} else {
												v387 = v16 + int32(16)
												v389 = v374 + int32(127)
												if v389&int32(64) != 0 {
													v408 = int64(0)
													v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
												} else {
													if v389 == int32(0) {
														v408 = v371
														v409 = v373
													} else {
														v400 = base.I64_extend_i32_u(v389)
														v408 = v371 << (uint(v400) % 64)
														v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
												*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
												v414 = int32(1) - v374
												if v414&int32(64) != 0 {
													v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
													v434 = int64(0)
												} else {
													if v414 == int32(0) {
														v433 = v371
														v434 = v373
													} else {
														v429 = base.I64_extend_i32_u(v414)
														v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
														v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
												v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
												v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
												v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
												v448 = v446
												v449 = v383
											}
											v452 = int64(3)
											v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
											v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
											v466 = base.I32_wrap_i64(v447) & int32(7)
											if v466 != int32(4) {
												v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
												v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
												if v466 == int32(0) {
													v488 = v472
													v489 = v475
												} else {
													v484 = v472
													v485 = v475
													v488 = v484
													v489 = v485
												}
											} else {
												v478 = v454 + v454&int64(1)
												v484 = v478
												v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
												v488 = v484
												v489 = v485
											}
										}
									}
								} else {
									v488 = l1
									v489 = l2
								}
							}
						}
					}
				}
			}
		}
	} else {
		v43 = int64(9223090561878065152)
		if v23 == v43 {
			v47 = v21
		} else {
			v47 = base.B2i32(base.Ui64(v23) < base.Ui64(v43))
		}
		if v47 == int32(0) {
			v488 = l1
			v489 = l2 | int64(140737488355328)
		} else {
			v54 = int64(9223090561878065152)
			if v19 == v54 {
				v58 = base.B2i32(l3 == int64(0))
			} else {
				v58 = base.B2i32(base.Ui64(v19) < base.Ui64(v54))
			}
			if v58 == int32(0) {
				v488 = l3
				v489 = l4 | int64(140737488355328)
			} else {
				if l1|(v23^int64(9223090561878065152)) == int64(0) {
					v75 = base.B2i32(l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0))
					if l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0) {
						v76 = int64(9223231299366420480)
					} else {
						v76 = l2
					}
					if l1^l3|(l2^l4^int64(-9223372036854775807-1)) == int64(0) {
						v78 = int64(0)
					} else {
						v78 = l1
					}
					v488 = v78
					v489 = v76
				} else {
					if l3|(v19^int64(9223090561878065152)) == int64(0) {
						v488 = l3
						v489 = l4
					} else {
						if l1|v23 == int64(0) {
							if l3|v19 != int64(0) {
								v488 = l3
								v489 = l4
							} else {
								v488 = l1 & l3
								v489 = l2 & l4
							}
						} else {
							if l3|v19 != int64(0) {
								if v19 == v23 {
									v99 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
								} else {
									v99 = base.B2i32(base.Ui64(v23) < base.Ui64(v19))
								}
								if v99 != 0 {
									v100 = l3
								} else {
									v100 = l1
								}
								if v99 != 0 {
									v101 = l4
								} else {
									v101 = l2
								}
								v103 = v101 & int64(281474976710655)
								if v99 != 0 {
									v104 = l2
								} else {
									v104 = l4
								}
								v105 = int64(48)
								v108 = int32(_a_F___addtf3_0)
								v109 = base.I32_wrap_i64(int64(base.Ui64(v104)>>(uint(v105)%64))) & v108
								v114 = base.I32_wrap_i64(int64(base.Ui64(v101)>>(uint(v105)%64))) & v108
								if v114 == int32(0) {
									v118 = v16 + int32(96)
									v120 = base.B2i32(v103 == int64(0))
									if v103 == int64(0) {
										v121 = v100
									} else {
										v121 = v103
									}
									if v103 == int64(0) {
										v125 = int64(64)
									} else {
										v125 = int64(0)
									}
									v127 = base.I32_wrap_i64(base.I64_clz(v121) + v125)
									v129 = v127 - int32(15)
									if v129&int32(64) != 0 {
										v148 = int64(0)
										v149 = v100 << (uint(base.I64_extend_i32_u(v129+int32(-64))) % 64)
									} else {
										if v129 == int32(0) {
											v148 = v100
											v149 = v103
										} else {
											v140 = base.I64_extend_i32_u(v129)
											v148 = v100 << (uint(v140) % 64)
											v149 = v103<<(uint(v140)%64) | int64(base.Ui64(v100)>>(uint(base.I64_extend_i32_u(int32(64)-v129))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v118))) = v148
									*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v149
									v153 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
									v157 = v154
									v158 = v153
									v159 = int32(16) - v127
								} else {
									v157 = v103
									v158 = v100
									v159 = v114
								}
								if v99 != 0 {
									v160 = l1
								} else {
									v160 = l3
								}
								v162 = v104 & int64(281474976710655)
								if v109 != 0 {
									v203 = v160
									v204 = v109
									v205 = v162
								} else {
									v164 = v16 + int32(80)
									v166 = base.B2i32(v162 == int64(0))
									if v162 == int64(0) {
										v167 = v160
									} else {
										v167 = v162
									}
									if v162 == int64(0) {
										v171 = int64(64)
									} else {
										v171 = int64(0)
									}
									v173 = base.I32_wrap_i64(base.I64_clz(v167) + v171)
									v175 = v173 - int32(15)
									if v175&int32(64) != 0 {
										v194 = int64(0)
										v195 = v160 << (uint(base.I64_extend_i32_u(v175+int32(-64))) % 64)
									} else {
										if v175 == int32(0) {
											v194 = v160
											v195 = v162
										} else {
											v186 = base.I64_extend_i32_u(v175)
											v194 = v160 << (uint(v186) % 64)
											v195 = v162<<(uint(v186)%64) | int64(base.Ui64(v160)>>(uint(base.I64_extend_i32_u(int32(64)-v175))%64))
										}
									}
									*(*int64)(unsafe.Add(mBase, uint32(v164))) = v194
									*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = v195
									v201 = *(*int64)(unsafe.Add(mBase, uint32(v16)+80))
									v202 = *(*int64)(unsafe.Add(mBase, uint32(v16)+88))
									v203 = v201
									v204 = int32(16) - v173
									v205 = v202
								}
								v206 = int64(3)
								v208 = int64(61)
								v212 = v205<<(uint(v206)%64) | int64(base.Ui64(v203)>>(uint(v208)%64)) | int64(2251799813685248)
								v220 = v203 << (uint(v206) % 64)
								if v159 == v204 {
									v288 = v212
									v290 = v220
								} else {
									v222 = v159 - v204
									if base.Ui32(int32(127)) < base.Ui32(v222) {
										v288 = int64(0)
										v290 = int64(1)
									} else {
										v228 = v16 - int32(-64)
										v230 = int32(128) - v222
										if v230&int32(64) != 0 {
											v249 = int64(0)
											v250 = v220 << (uint(base.I64_extend_i32_u(v230+int32(-64))) % 64)
										} else {
											if v230 == int32(0) {
												v249 = v220
												v250 = v212
											} else {
												v241 = base.I64_extend_i32_u(v230)
												v249 = v220 << (uint(v241) % 64)
												v250 = v212<<(uint(v241)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v230))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v228))) = v249
										*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = v250
										v255 = v16 + int32(48)
										if v222&int32(64) != 0 {
											v274 = int64(base.Ui64(v212) >> (uint(base.I64_extend_i32_u(v222+int32(-64))) % 64))
											v275 = int64(0)
										} else {
											if v222 == int32(0) {
												v274 = v220
												v275 = v212
											} else {
												v270 = base.I64_extend_i32_u(v222)
												v274 = v212<<(uint(base.I64_extend_i32_u(int32(64)-v222))%64) | int64(base.Ui64(v220)>>(uint(v270)%64))
												v275 = int64(base.Ui64(v212) >> (uint(v270) % 64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v255))) = v274
										*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v275
										v279 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
										v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
										v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
										v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
										v288 = v279
										v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
									}
								}
								v292 = v157<<(uint(v206)%64) | int64(base.Ui64(v158)>>(uint(v208)%64)) | int64(2251799813685248)
								v294 = v158 << (uint(int64(3)) % 64)
								if l2^l4 < int64(0) {
									v297 = int64(0)
									if v290^v294|(v288^v292) == v297 {
										v488 = v297
										v489 = v297
									} else {
										v304 = v294 - v290
										v308 = v292 - v288 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v294) < base.Ui64(v290)))
										if base.Ui64(int64(2251799813685247)) < base.Ui64(v308) {
											v371 = v304
											v373 = v308
											v374 = v159
										} else {
											v312 = v16 + int32(32)
											v314 = base.B2i32(v308 == int64(0))
											if v308 == int64(0) {
												v315 = v304
											} else {
												v315 = v308
											}
											if v308 == int64(0) {
												v319 = int64(64)
											} else {
												v319 = int64(0)
											}
											v323 = base.I32_wrap_i64(base.I64_clz(v315)|v319) - int32(12)
											if v323&int32(64) != 0 {
												v342 = int64(0)
												v343 = v304 << (uint(base.I64_extend_i32_u(v323+int32(-64))) % 64)
											} else {
												if v323 == int32(0) {
													v342 = v304
													v343 = v308
												} else {
													v334 = base.I64_extend_i32_u(v323)
													v342 = v304 << (uint(v334) % 64)
													v343 = v308<<(uint(v334)%64) | int64(base.Ui64(v304)>>(uint(base.I64_extend_i32_u(int32(64)-v323))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v312))) = v342
											*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = v343
											v348 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
											v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
											v371 = v349
											v373 = v348
											v374 = v159 - v323
										}
										v377 = v101 & int64(-9223372036854775807-1)
										if int32(_a_F___addtf3_0) <= v374 {
											v488 = int64(0)
											v489 = v377 | int64(9223090561878065152)
										} else {
											v383 = int32(0)
											if v383 < v374 {
												v447 = v371
												v448 = v373
												v449 = v374
											} else {
												v387 = v16 + int32(16)
												v389 = v374 + int32(127)
												if v389&int32(64) != 0 {
													v408 = int64(0)
													v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
												} else {
													if v389 == int32(0) {
														v408 = v371
														v409 = v373
													} else {
														v400 = base.I64_extend_i32_u(v389)
														v408 = v371 << (uint(v400) % 64)
														v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
												*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
												v414 = int32(1) - v374
												if v414&int32(64) != 0 {
													v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
													v434 = int64(0)
												} else {
													if v414 == int32(0) {
														v433 = v371
														v434 = v373
													} else {
														v429 = base.I64_extend_i32_u(v414)
														v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
														v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
												v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
												v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
												v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
												v448 = v446
												v449 = v383
											}
											v452 = int64(3)
											v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
											v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
											v466 = base.I32_wrap_i64(v447) & int32(7)
											if v466 != int32(4) {
												v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
												v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
												if v466 == int32(0) {
													v488 = v472
													v489 = v475
												} else {
													v484 = v472
													v485 = v475
													v488 = v484
													v489 = v485
												}
											} else {
												v478 = v454 + v454&int64(1)
												v484 = v478
												v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
												v488 = v484
												v489 = v485
											}
										}
									}
								} else {
									v350 = v290 + v294
									v354 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v350) < base.Ui64(v290))) + (v288 + v292)
									if v354&int64(4503599627370496) == int64(0) {
										v371 = v350
										v373 = v354
										v374 = v159
									} else {
										v359 = int64(1)
										v371 = v290&v359 | (v354<<(uint(int64(63))%64) | int64(base.Ui64(v350)>>(uint(v359)%64)))
										v373 = int64(base.Ui64(v354) >> (uint(v359) % 64))
										v374 = v159 + int32(1)
									}
									v377 = v101 & int64(-9223372036854775807-1)
									if int32(_a_F___addtf3_0) <= v374 {
										v488 = int64(0)
										v489 = v377 | int64(9223090561878065152)
									} else {
										v383 = int32(0)
										if v383 < v374 {
											v447 = v371
											v448 = v373
											v449 = v374
										} else {
											v387 = v16 + int32(16)
											v389 = v374 + int32(127)
											if v389&int32(64) != 0 {
												v408 = int64(0)
												v409 = v371 << (uint(base.I64_extend_i32_u(v374+int32(63))) % 64)
											} else {
												if v389 == int32(0) {
													v408 = v371
													v409 = v373
												} else {
													v400 = base.I64_extend_i32_u(v389)
													v408 = v371 << (uint(v400) % 64)
													v409 = v373<<(uint(v400)%64) | int64(base.Ui64(v371)>>(uint(base.I64_extend_i32_u(int32(64)-v389))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v387))) = v408
											*(*int64)(unsafe.Add(mBase, uint32(v387)+8)) = v409
											v414 = int32(1) - v374
											if v414&int32(64) != 0 {
												v433 = int64(base.Ui64(v373) >> (uint(base.I64_extend_i32_u(v414+int32(-64))) % 64))
												v434 = int64(0)
											} else {
												if v414 == int32(0) {
													v433 = v371
													v434 = v373
												} else {
													v429 = base.I64_extend_i32_u(v414)
													v433 = v373<<(uint(base.I64_extend_i32_u(int32(64)-v414))%64) | int64(base.Ui64(v371)>>(uint(v429)%64))
													v434 = int64(base.Ui64(v373) >> (uint(v429) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v16))) = v433
											*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v434
											v438 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
											v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
											v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
											v446 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
											v447 = v438 | base.I64_extend_i32_u(base.B2i32(v439|v440 != int64(0)))
											v448 = v446
											v449 = v383
										}
										v452 = int64(3)
										v454 = v448<<(uint(int64(61))%64) | int64(base.Ui64(v447)>>(uint(v452)%64))
										v463 = int64(base.Ui64(v448)>>(uint(v452)%64))&int64(281474976710655) | base.I64_extend_i32_u(v449)<<(uint(int64(48))%64) | v377
										v466 = base.I32_wrap_i64(v447) & int32(7)
										if v466 != int32(4) {
											v472 = v454 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v466)))
											v475 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v472) < base.Ui64(v454)))
											if v466 == int32(0) {
												v488 = v472
												v489 = v475
											} else {
												v484 = v472
												v485 = v475
												v488 = v484
												v489 = v485
											}
										} else {
											v478 = v454 + v454&int64(1)
											v484 = v478
											v485 = v463 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v478) < base.Ui64(v454)))
											v488 = v484
											v489 = v485
										}
									}
								}
							} else {
								v488 = l1
								v489 = l2
							}
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v488
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v489
	m.G0 = v16 + int32(112)
	return
}
func F_aclequal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	v3 = int32(0)
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		v106 = v3
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return int32(1)
L7:
	;
	goto L8
L8:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	return base.B2i32(v11 == int32(0))
L9:
	;
	return v106
L10:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v5 != v17 {
		v106 = v3
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v27 = v19
	goto L14
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = (v20<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L14
L14:
	;
	v28 = v27 + l0
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = v29
	goto L17
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v38 = v37 + l1
	v39 = int32(4)
	v40 = v5 << (uint(v39) % 32)
	if base.Ui32(v39) <= base.Ui32(v40) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v106 = base.B2i32(v102 == int32(0))
	goto L9
L19:
	;
	v102 = int32(0)
	goto L18
L20:
	;
	v76 = v71
	v77 = v72
	v78 = v73
	goto L30
L21:
	;
	if (v28|v38)&int32(3) != 0 {
		v71 = v28
		v72 = v38
		v73 = v40
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v64 = v28
	v65 = v38
	v66 = v40
	goto L23
L23:
	;
	if v66 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L24:
	;
	v48 = v28
	v49 = v38
	v50 = v40
	goto L25
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v53 != v54 {
		v71 = v48
		v72 = v49
		v73 = v50
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v64 = v59
	v65 = v57
	v66 = v61
	goto L23
L27:
	;
	v56 = int32(4)
	v57 = v49 + v56
	v59 = v48 + v56
	v61 = v50 - v56
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		v48 = v59
		v49 = v57
		v50 = v61
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v71 = v64
	v72 = v65
	v73 = v66
	goto L20
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 == v82 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v102 = v81 - v82
	goto L18
L32:
	;
	v84 = int32(1)
	v89 = v78 - v84
	if v89 != 0 {
		v76 = v76 + v84
		v77 = v77 + v84
		v78 = v89
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L19
}
func F_aclitem_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 != v7 {
		v16 = v3
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v9 != v10 {
			v16 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v16 = base.I64_extend_i32_u(base.B2i32(v12 == v13))
		}
	}
	return v16
}
func F_aclmask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v224 int64
	_ = v224
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	v6 = int64(0)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_check_acl(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L83
	}
L4:
	;
	return int64(0)
L5:
	;
	if l3 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int64(0)
L7:
	;
	goto L8
L8:
	;
	v21 = l3 & int64(-4294967296)
	if v21 == int64(0) {
		v74 = v6
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v224
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	if l1 == l2 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l4 != 0 {
		v224 = v21
		goto L9
	} else {
		goto L31
	}
L13:
	;
	v25 = F_superuser_arg(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v25 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v28 = int32(0)
	v30 = F_roles_is_member_of(m, l1, int32(1), v28, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v32 = int32(0)
	if v30 == v32 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v70 == int32(0) {
		v74 = v6
		goto L10
	} else {
		goto L30
	}
L18:
	;
	v70 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v38 <= int32(0) {
		v64 = v32
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = v64
	goto L17
L22:
	;
	v41 = int32(0)
	if v41 < v38 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v44 = v38
	goto L25
L24:
	;
	v44 = v41
	goto L25
L25:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v47 = int32(0)
	goto L26
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45+v47<<(uint(int32(2))%32))))
	v56 = base.B2i32(v55 == l2)
	if v55 == l2 {
		v64 = v56
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v64 = v56
	goto L21
L28:
	;
	v58 = v47 + int32(1)
	if v58 != v44 {
		v47 = v58
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L12
L31:
	;
	if l3 == v21 {
		v224 = v21
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v74 = v21
	goto L10
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = (v79<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L35
L34:
	;
	v86 = v76
	goto L35
L35:
	;
	if v75 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return v74
L37:
	;
	goto L38
L38:
	;
	v90 = l0 + v86
	v92 = int32(0)
	v98 = v74
	goto L39
L39:
	;
	v105 = v90 + v92<<(uint(int32(4))%32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if l1 != v106 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v129 = int32(0)
	v134 = v121
	v136 = l3 & (v121 ^ int64(-1))
	goto L52
L41:
	;
	v123 = v92 + int32(1)
	if v123 != v75 {
		v92 = v123
		v98 = v121
		goto L39
	} else {
		goto L51
	}
L42:
	;
	v109 = v106
	goto L44
L43:
	;
	v109 = int32(0)
	goto L44
L44:
	;
	if v109 != 0 {
		v121 = v98
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v105)+8))
	v112 = v110&l3 | v98
	if l4 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v112 != l3 {
		v121 = v112
		goto L41
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v117 = int64(0)
	if v112 != v117 {
		v224 = v112
		goto L9
	} else {
		goto L50
	}
L49:
	;
	return l3
L50:
	;
	v121 = v117
	goto L41
L51:
	;
	goto L40
L52:
	;
	v142 = v90 + v129<<(uint(int32(4))%32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if base.B2i32(v143 == int32(0))|base.B2i32(l1 == v143) != 0 {
		v214 = v134
		v215 = v136
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v224 = v214
	goto L9
L54:
	;
	v217 = v129 + int32(1)
	if v217 != v75 {
		v129 = v217
		v134 = v214
		v136 = v215
		goto L52
	} else {
		goto L82
	}
L55:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	if v148&v136 == int64(0) {
		v214 = v134
		v215 = v136
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v152 = F_superuser_arg(m, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v152 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v157 = int32(0)
	v159 = F_roles_is_member_of(m, l1, int32(1), v157, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	v204 = v202&l3 | v134
	if l4 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L61:
	;
	v161 = int32(0)
	if v159 == v161 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v199 == int32(0) {
		v214 = v134
		v215 = v136
		goto L54
	} else {
		goto L75
	}
L63:
	;
	v199 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v167 <= int32(0) {
		v193 = v161
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = v193
	goto L62
L67:
	;
	v170 = int32(0)
	if v170 < v167 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v173 = v167
	goto L70
L69:
	;
	v173 = v170
	goto L70
L70:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v176 = int32(0)
	goto L71
L71:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v174+v176<<(uint(int32(2))%32))))
	v185 = base.B2i32(v184 == v143)
	if v184 == v143 {
		v193 = v185
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v193 = v185
	goto L66
L73:
	;
	v187 = v176 + int32(1)
	if v187 != v173 {
		v176 = v187
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L60
L76:
	;
	v214 = v204
	v215 = l3 & (v204 ^ int64(-1))
	goto L54
L77:
	;
	if l3 != v204 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v204 != int64(0) {
		v224 = v204
		goto L9
	} else {
		goto L81
	}
L80:
	;
	return l3
L81:
	;
	goto L76
L82:
	;
	goto L53
L83:
	;
	F_errmsg_internal(m, int32(_a_F_aclmask_0), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_aclmask_1), int32(1426), int32(_a_F_aclmask_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_aclremove(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errmsg(m, int32(_a_F_aclremove_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_aclremove_1), int32(1630), int32(_a_F_aclremove_2))
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_action_terminate(m *base.Module, l0 int32) {
	m.Env.X_emscripten_runtime_keepalive_clear(m)
	F__Exit(m, l0+int32(128))
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addKey(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v89 = F_lappend(m, v85, l2)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L26
	} else {
		goto L29
	}
L3:
	;
	v16 = int32(0)
	v18 = v11
	goto L6
L4:
	;
	goto L5
L5:
	;
	v85 = int32(0)
	goto L2
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v16 {
		v85 = v18
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v16<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v29 != v30 {
		v61 = v16
		v63 = v18
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v63 != 0 {
		v16 = v61 + int32(1)
		v18 = v63
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v32 == int32(-3) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v35 != int32(-3) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v45 == int32(-3) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.B2i32(v35 != v38)|base.B2i32(v32 != v40) != 0 {
		v45 = v40
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 == v43 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L1
L17:
	;
	v45 = v43
	goto L12
L18:
	;
	v55 = F_list_delete_nth_cell(m, v18, v16)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v48 != int32(-3) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v45 != v32 {
		v61 = v16
		v63 = v18
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v45 != v32 {
		v61 = v16
		v63 = v18
		goto L9
	} else {
		goto L25
	}
L23:
	;
	if v48 == v35 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v61 = v16
	v63 = v18
	goto L9
L25:
	;
	goto L18
L26:
	;
	return
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v55
	v61 = v16 - int32(1)
	v63 = v55
	goto L9
L28:
	;
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v89
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+24))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+36))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v94 == v95 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v97 | int32(2)
	return
L31:
	;
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = F_pg_reg_getnumoutarcs(m, v102, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v106 = F_palloc_mul(m, int32(8), v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_pg_reg_getoutarcs(m, v108, v109, v106, v104)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	if int32(0) < v104 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v121 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v106)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L26
	} else {
		goto L75
	}
L39:
	;
	v125 = int32(-4)
	v129 = v106 + v121<<(uint(int32(3))%32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
	v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+40)))
	if v130 != v133 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L38
L41:
	;
	v222 = v121 + int32(1)
	if v222 != v104 {
		v121 = v222
		goto L39
	} else {
		goto L74
	}
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v208 = F_palloc(m, int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L72
	}
L43:
	;
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+42)))
	v138 = base.B2i32(v130 == v135)
	goto L45
L44:
	;
	v138 = int32(1)
	goto L45
L45:
	;
	if v138 != 0 {
		v203 = v125
		v204 = v125
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v139 = int32(-3)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+44)))
	if v140 != v143 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+46)))
	v148 = base.B2i32(v140 == v145)
	goto L49
L48:
	;
	v148 = int32(1)
	goto L49
L49:
	;
	if v148 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v203 = int32(-3)
	v204 = v139
	goto L42
L51:
	;
	goto L52
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v150 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v203 = int32(-3)
	v204 = v139
	goto L42
L54:
	;
	goto L55
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = v155 + v150*int32(12)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != int32(1) {
		v203 = int32(-3)
		v204 = v139
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v162 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v187 <= int32(0) {
		goto L41
	} else {
		goto L66
	}
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v165 + int32(4) {
	case 0:
		goto L61
	case 1:
		goto L59
	default:
		goto L60
	}
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v176 = F_palloc(m, int32(12))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L26
	} else {
		goto L64
	}
L60:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v171 != int32(-4) {
		goto L57
	} else {
		goto L63
	}
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v168 == int32(-4) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	goto L59
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v174
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = int64(-12884901892)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v182 = F_lappend(m, v181, v176)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L26
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v182
	goto L57
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v192 + int32(4) {
	case 0:
		goto L68
	case 1:
		v203 = v191
		v204 = v190
		goto L42
	default:
		goto L67
	}
L67:
	;
	v200 = int32(-4)
	if v190 != v200 {
		goto L41
	} else {
		goto L71
	}
L68:
	;
	if v191 != int32(-4) {
		goto L41
	} else {
		goto L69
	}
L69:
	;
	v197 = int32(-4)
	if v190 == v197 {
		v203 = v191
		v204 = v197
		goto L42
	} else {
		goto L70
	}
L70:
	;
	goto L41
L71:
	;
	v203 = v191
	v204 = v200
	goto L42
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v204
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v214 = F_lappend(m, v213, v208)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L26
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v214
	goto L41
L74:
	;
	goto L40
L75:
	;
	goto L1
}
func F_addOrReplaceTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v11) {
		v19 = int32(base.Ui32(v11+int32(_a_F_addOrReplaceTuple_0)) >> (uint(int32(2)) % 32))
	} else {
		v19 = int32(0)
	}
	if base.Ui32(l3) <= base.Ui32(v19&int32(_a_F_addOrReplaceTuple_1)) {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+l3<<(uint(int32(2))%32))+20))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v26&int32(_a_F_addOrReplaceTuple_2))))
		v31 = int32(3)
		if v30&v31 != v31 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_addOrReplaceTuple_3), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_addOrReplaceTuple_4), int32(58), int32(_a_F_addOrReplaceTuple_5))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			v36 = l0 + v35
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
			v39 = v37 - int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v39)
			F_PageIndexTupleDelete(m, l0, l3)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v45 = F_PageAddItemExtended(m, l0, l1, l2, l3, int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					if v45 != l3 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg_internal(m, int32(_a_F_addOrReplaceTuple_6), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_addOrReplaceTuple_4), int32(70), int32(_a_F_addOrReplaceTuple_5))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v45 = F_PageAddItemExtended(m, l0, l1, l2, l3, int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			if v45 != l3 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
					F_errmsg_internal(m, int32(_a_F_addOrReplaceTuple_6), v9)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_addOrReplaceTuple_4), int32(70), int32(_a_F_addOrReplaceTuple_5))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_add_abs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = int32(-1)
	v26 = v22 + (v23 ^ v24)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = v27 + (v28 ^ v24)
	if v31 < v26 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v26
	goto L3
L2:
	;
	v33 = v31
	goto L3
L3:
	;
	v35 = v33 + int32(1)
	if v28 < v23 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = v23
	goto L6
L5:
	;
	v37 = v28
	goto L6
L6:
	;
	v38 = int32(1)
	v39 = v37 + v38
	v40 = v35 + v39
	if v40 <= v38 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = int32(1)
	goto L9
L8:
	;
	v43 = v40
	goto L9
L9:
	;
	v45 = v43 << (uint(int32(1)) % 32)
	v48 = F_palloc(m, v45+int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v50 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v50)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = v43
	v61 = v50
	v62 = v35 + v54
	v63 = v35 + v52
	goto L12
L12:
	;
	v73 = int32(1)
	v79 = v62 - v73
	v80 = int32(0)
	if base.B2i32(v79 < v80)|base.B2i32(v22 <= v79) == v80 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v111 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v106 = base.B2i32(int32(_a_F_add_abs_0) < v102)
	if int32(_a_F_add_abs_0) < v102 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v79<<(uint(int32(1))%32)))))
	v91 = v61 + v89
	goto L17
L16:
	;
	v91 = v61
	goto L17
L17:
	;
	v93 = v63 - int32(1)
	if v93 < int32(0) {
		v102 = v91
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v27 <= v93 {
		v102 = v91
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v93<<(uint(int32(1))%32)))))
	v102 = v91 + v100
	goto L14
L20:
	;
	v107 = v102 - int32(_a_F_add_abs_1)
	goto L22
L21:
	;
	v107 = v102
	goto L22
L22:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48+v57<<(uint(v73)%32)))) = uint16(v107)
	if base.Ui32(int32(1)) < base.Ui32(v57) {
		v57 = v57 - v73
		v61 = v106
		v62 = v79
		v63 = v93
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L13
L24:
	;
	F_pfree(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
	if v17 < v18 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v117 = v18
	goto L30
L29:
	;
	v117 = v17
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v39
	v121 = v48 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v121
	v125 = v121
	v127 = v43
	v131 = v39
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v198
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v196
	return
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v196 = v177
	v198 = int32(0)
	goto L31
L33:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125))))
	if v140 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v177 = v121 + v45
	goto L32
L35:
	;
	v144 = v127
	goto L38
L36:
	;
	goto L37
L37:
	;
	v167 = int32(1)
	v168 = v131 - v167
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v168
	if v167 < v127 {
		v125 = v125 + int32(2)
		v127 = v127 - v167
		v131 = v168
		goto L33
	} else {
		goto L42
	}
L38:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125+v144<<(uint(int32(1))%32)-int32(2)))))
	if v162 != 0 {
		v196 = v125
		v198 = v144
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v163 = int32(1)
	if v163 < v144 {
		v144 = v144 - v163
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v177 = v125
	goto L32
L42:
	;
	goto L34
}
func F_add_vars_to_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v12 + int32(16)
	return
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = v4
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v135 = v26 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v135 < v136 {
		v26 = v135
		goto L5
	} else {
		goto L37
	}
L8:
	;
	if v33 != int32(321) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v45 = F_find_base_rel(m, l0, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v38 = F_find_placeholder_info(m, l0, v32)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v41 = F_bms_add_members(m, v40, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v41
	goto L7
L15:
	;
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+8)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v49 = int32(0)
	if l2 == v49 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v102 != 0 {
		goto L7
	} else {
		goto L30
	}
L17:
	;
	v102 = int32(1)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v48 == int32(0) {
		v95 = v49
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v102 = v95
	goto L16
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v59 < v58 {
		v95 = v49
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v61 = int32(1)
	if v58 <= v61 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v64 = v61
	goto L25
L24:
	;
	v64 = v58
	goto L25
L25:
	;
	v65 = int32(8)
	v70 = int32(0)
	goto L26
L26:
	;
	v77 = v70 << (uint(int32(2)) % 32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2+v65+v77)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48+v65+v77)))
	v84 = v79 & (v81 ^ int32(-1))
	v86 = base.B2i32(v84 == int32(0))
	if v84 != 0 {
		v95 = v86
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v95 = v86
	goto L20
L28:
	;
	v88 = v70 + int32(1)
	if v88 != v64 {
		v70 = v88
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+88)))
	v106 = (v47 - v103) << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v45)+92))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106+v107)))
	if v109 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v124 = v109
	goto L33
L32:
	;
	v110 = F_copyObjectImpl(m, v32)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v125 = F_bms_add_members(m, v124, l2)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v116 = F_lappend(m, v115, v110)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v45)+92))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120+v106)))
	v124 = v122
	goto L33
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v45)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v106))) = v125
	goto L7
L37:
	;
	goto L6
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v154
	F_errmsg_internal(m, int32(_a_F_add_vars_to_targetlist_0), v12)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_add_vars_to_targetlist_1), int32(345), int32(_a_F_add_vars_to_targetlist_2))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_adjacent_inner_consistent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l3 != 0 {
		v12 = F_range_cmp_bounds(m, l0, l1, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)))
			if v16 == int32(1) {
				if v12 < int32(0) {
					v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v22
					v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v24
					v26 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v26
					v28 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v28
					v34 = F_bounds_adjacent(m, l0, v8+int32(-16), v8+int32(-32))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 == int32(0) {
							v44 = int32(-1)
						} else {
							v44 = int32(1)
						}
						v46 = F_range_cmp_bounds(m, l0, l2, l3)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = int32(0)
							v49 = base.B2i32(v48 <= v46)
							if v49&base.B2i32(v44 < v48)|base.B2i32(base.B2i32(v44 <= v48)|v49 == v48) != 0 {
								v91 = int32(0)
								m.G0 = v10 - int32(-64)
								return v91
							} else {
								v62 = F_range_cmp_bounds(m, l0, l1, l2)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
									if v64 == int32(1) {
										if v62 < int32(0) {
											v70 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v70
											v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v72
											v74 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v74
											v76 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v76
											v80 = F_bounds_adjacent(m, l0, v8+int32(-48), v10)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												if v80 == int32(0) {
													v91 = int32(-1)
												} else {
													v91 = int32(1)
												}
												m.G0 = v10 - int32(-64)
												return v91
											}
										} else {
											v91 = int32(1)
											m.G0 = v10 - int32(-64)
											return v91
										}
									} else {
										if v62 <= int32(0) {
											v89 = int32(-1)
										} else {
											v89 = int32(1)
										}
										v91 = v89
										m.G0 = v10 - int32(-64)
										return v91
									}
								}
							}
						}
					}
				} else {
					v44 = int32(1)
					v46 = F_range_cmp_bounds(m, l0, l2, l3)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = int32(0)
						v49 = base.B2i32(v48 <= v46)
						if v49&base.B2i32(v44 < v48)|base.B2i32(base.B2i32(v44 <= v48)|v49 == v48) != 0 {
							v91 = int32(0)
							m.G0 = v10 - int32(-64)
							return v91
						} else {
							v62 = F_range_cmp_bounds(m, l0, l1, l2)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
								if v64 == int32(1) {
									if v62 < int32(0) {
										v70 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v70
										v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v72
										v74 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v76
										v80 = F_bounds_adjacent(m, l0, v8+int32(-48), v10)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											if v80 == int32(0) {
												v91 = int32(-1)
											} else {
												v91 = int32(1)
											}
											m.G0 = v10 - int32(-64)
											return v91
										}
									} else {
										v91 = int32(1)
										m.G0 = v10 - int32(-64)
										return v91
									}
								} else {
									if v62 <= int32(0) {
										v89 = int32(-1)
									} else {
										v89 = int32(1)
									}
									v91 = v89
									m.G0 = v10 - int32(-64)
									return v91
								}
							}
						}
					}
				}
			} else {
				if v12 <= int32(0) {
					v43 = int32(-1)
				} else {
					v43 = int32(1)
				}
				v44 = v43
				v46 = F_range_cmp_bounds(m, l0, l2, l3)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = int32(0)
					v49 = base.B2i32(v48 <= v46)
					if v49&base.B2i32(v44 < v48)|base.B2i32(base.B2i32(v44 <= v48)|v49 == v48) != 0 {
						v91 = int32(0)
						m.G0 = v10 - int32(-64)
						return v91
					} else {
						v62 = F_range_cmp_bounds(m, l0, l1, l2)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
							if v64 == int32(1) {
								if v62 < int32(0) {
									v70 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v70
									v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v72
									v74 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v76
									v80 = F_bounds_adjacent(m, l0, v8+int32(-48), v10)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										if v80 == int32(0) {
											v91 = int32(-1)
										} else {
											v91 = int32(1)
										}
										m.G0 = v10 - int32(-64)
										return v91
									}
								} else {
									v91 = int32(1)
									m.G0 = v10 - int32(-64)
									return v91
								}
							} else {
								if v62 <= int32(0) {
									v89 = int32(-1)
								} else {
									v89 = int32(1)
								}
								v91 = v89
								m.G0 = v10 - int32(-64)
								return v91
							}
						}
					}
				}
			}
		}
	} else {
		v62 = F_range_cmp_bounds(m, l0, l1, l2)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)))
			if v64 == int32(1) {
				if v62 < int32(0) {
					v70 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v70
					v72 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v72
					v74 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v74
					v76 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v76
					v80 = F_bounds_adjacent(m, l0, v8+int32(-48), v10)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						if v80 == int32(0) {
							v91 = int32(-1)
						} else {
							v91 = int32(1)
						}
						m.G0 = v10 - int32(-64)
						return v91
					}
				} else {
					v91 = int32(1)
					m.G0 = v10 - int32(-64)
					return v91
				}
			} else {
				if v62 <= int32(0) {
					v89 = int32(-1)
				} else {
					v89 = int32(1)
				}
				v91 = v89
				m.G0 = v10 - int32(-64)
				return v91
			}
		}
	}
}
func F_adjust_child_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v4 = int32(0)
	if v4 < l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v4
	v13 = v4
	goto L4
L2:
	;
	v40 = v4
	goto L3
L3:
	;
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2+v13<<(uint(int32(2))%32))))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = F_bms_is_member(m, v19, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v40 = v33
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v12 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v33 = v12
	goto L10
L10:
	;
	v35 = v13 + int32(1)
	if v35 != l1 {
		v12 = v33
		v13 = v35
		goto L4
	} else {
		goto L17
	}
L11:
	;
	v26 = v12
	goto L13
L12:
	;
	v24 = F_bms_copy(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v28 = F_bms_del_member(m, v26, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	v26 = v24
	goto L13
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v31 = F_bms_add_member(m, v28, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v33 = v31
	goto L10
L17:
	;
	goto L5
L18:
	;
	v43 = v40
	goto L20
L19:
	;
	v43 = l0
	goto L20
L20:
	;
	return v43
}
func F_af_6_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v5 - int32(1) {
	case 0:
		v8 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 <= v10 {
			v87 = v8
			return v87
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9-int32(1)))))
			if v16 != int32(105) {
				v87 = v8
				return v87
			} else {
				v20 = v9 - int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v20 <= v31 {
					v74 = int32(-1)
				} else {
					v43 = int32(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v20-v43))))
					if int32(246) < v48 {
						v70 = v43
					} else {
						v50 = v48 - int32(97)
						if v50 < int32(0) {
							v70 = v43
						} else {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v50)>>(uint(int32(3))%32)))+uint32(_c_F_af_6_1[0]))))
							if int32(base.Ui32(v56)>>(uint(v50&int32(7))%32))&int32(1) == int32(0) {
								v70 = v43
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - int32(1)
								v70 = int32(0)
							}
						}
					}
					v74 = v70
				}
				return base.B2i32(v74 == int32(0))
			}
		}
	case 1:
		v81 = F_find_among_b(m, l0, int32(_a_F_af_6_1_0), int32(7), int32(0))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return int32(0)
		} else {
			v87 = base.B2i32(v81 != int32(0))
			return v87
		}
	default:
		v87 = int32(-1)
		return v87
	}
}
func F_alloc_chromo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_palloc(m, int32(24))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v11 = F_palloc_mul(m, int32(4), l0+int32(1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v11
			return v4
		}
	}
}
func F_anyarray_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_anycompatiblearray_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatiblearray_recv_0), int32(175), int32(_a_F_anycompatiblearray_recv_1), int32(_a_F_anycompatiblearray_recv_2), int32(_a_F_anycompatiblearray_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14228(m, l0, l1, int32(_a_F_anytime_typmod_check_0), int32(71), int32(_a_F_anytime_typmod_check_1), int32(_a_F_anytime_typmod_check_2), int32(78), int32(_a_F_anytime_typmod_check_3))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_append_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v3 = int32(0)
	if l1 == v3 {
		v77 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v77
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		v77 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = l0
	v17 = v3
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v17<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+40)))
	if v28 != 0 {
		v65 = v14
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v77 = v65
	goto L1
L6:
	;
	v74 = v17 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 < v75 {
		v14 = v65
		v17 = v74
		goto L4
	} else {
		goto L17
	}
L7:
	;
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v61 = F_lappend(m, v14, v26)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v31 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v38 = int32(0)
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34+v38<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v27 == v48 {
		v65 = v14
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v51 = v38 + int32(1)
	if v31 != v51 {
		v38 = v51
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return int32(0)
L16:
	;
	v65 = v61
	goto L6
L17:
	;
	goto L5
}
func F_apply_tlist_labeling(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v10 = int32(0)
	goto L1
L1:
	;
	v11 = int32(0)
	if l0 == v11 {
		v21 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v10 {
		v21 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v17 + v10<<(uint(int32(2))%32)
	goto L3
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+v10<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v40
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+24)) = uint16(v42)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+26)) = uint8(v44)
	v10 = v10 + int32(1)
	goto L1
L7:
	;
	return
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(v21 == int32(0))|base.B2i32(v26 <= v10) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v29 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
func F_apw_dump_now(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(1200)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v17 = F_LWLockAcquire(m, v15, v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v23 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v335 = int32(_a_F_apw_dump_now_0)
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v338 = v12 + int32(176)
	v339 = F_unlink(m, v338)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v336
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L78
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L74
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L71
	}
L6:
	;
	m.G0 = v12 + int32(1200)
	return v291
L7:
	;
	F_LWLockRelease(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v51
	F_LWLockRelease(m, v22)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	if l0 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v32 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v32 == int32(0) {
		v291 = v3
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v38
	F_errmsg(m, int32(_a_F_apw_dump_now_1), v12+int32(144))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(701), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v291 = v3
	goto L6
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	v60 = F_palloc_extended(m, v56*int32(20), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	if int32(0) < v63 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = int32(0)
	v71 = v3
	goto L21
L19:
	;
	v130 = v3
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = int32(_a_F_apw_dump_now_4)
	v139 = v12 + int32(176)
	v144 = F_pg_snprintf(m, v139, int32(1024), int32(_a_F_apw_dump_now_5), v12+int32(128))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[4]))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v130 = v116
	goto L20
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[5]))
	v84 = v81 + v67*int32(56)
	v85 = F_LockBufHdr(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v89 = int64(0)
	v91 = int32(0)
	if base.B2i32(v85&int64(33554432) == v89)|base.B2i32(l1 == v91)&base.B2i32(v85&int64(2147483648) == v89) == v91 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v60 + v71*int32(20)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v112
	v116 = v71 + int32(1)
	goto L30
L29:
	;
	v116 = v71
	goto L30
L30:
	;
	v120 = base.AtomicRmwSub64(m, v84, int32(24), int64(4194304))
	v122 = v67 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	if v122 < v124 {
		v67 = v122
		v71 = v116
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L22
L32:
	;
	v147 = F_AllocateFile(m, v139, int32(_a_F_apw_dump_now_6))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v147 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v130
	v155 = F_pg_fprintf(m, v147, int32(_a_F_apw_dump_now_7), v12+int32(112))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	F_pfree(m, v60)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L63
	}
L36:
	;
	v190 = int32(0)
	goto L47
L37:
	;
	if int32(0) <= v155 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v130 <= int32(0) {
		goto L35
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v165 = F_FreeFile(m, v147)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L36
L42:
	;
	v168 = v12 + int32(176)
	v169 = F_unlink(m, v168)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v164
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v168
	F_errmsg(m, int32(_a_F_apw_dump_now_8), v12+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(766), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[4]))
	if v199 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v223 = F_FreeFile(m, v147)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L58
	}
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v204 = v60 + v190*int32(20)
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v204)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(96)))) = v207
	*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v206
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v205
	v214 = F_pg_fprintf(m, v147, int32(_a_F_apw_dump_now_9), v12+int32(80))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	if int32(0) <= v214 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v219 = v190 + int32(1)
	if v219 == v130 {
		goto L35
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L48
L57:
	;
	v190 = v219
	goto L47
L58:
	;
	v226 = v12 + int32(176)
	v227 = F_unlink(m, v226)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v222
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v226
	F_errmsg(m, int32(_a_F_apw_dump_now_8), v12+int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(789), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v258 = F_FreeFile(m, v147)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v258 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v264 = F_durable_rename(m, v12+int32(176), int32(_a_F_apw_dump_now_4), int32(21))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+20)) = int32(-1)
	v272 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v272 == int32(0) {
		v291 = v130
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v130
	F_errmsg_internal(m, int32(_a_F_apw_dump_now_10), v12+int32(48))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(816), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v291 = v130
	goto L6
L71:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v306
	F_errmsg(m, int32(_a_F_apw_dump_now_11), v12+int32(160))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(697), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(176)
	F_errmsg(m, int32(_a_F_apw_dump_now_12), v12)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(753), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v338
	F_errmsg(m, int32(_a_F_apw_dump_now_13), v12-int32(-64))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(809), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_apw_start_leader_worker(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	v8 = m.G0
	v10 = v8 - int32(1488)
	m.G0 = v10
	base.MemoryFill(m, v10+int32(32), int32(0), int32(1456))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+208)) = int64(4294967297)
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+220)) = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+227)) = v23
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1244)) = v26
	v29 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+1252)) = v29
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+1260)) = uint8(v32)
	v35 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v35
	v38 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+31)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v10)+127)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v35
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[8])))
	if v47 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L35
	}
L3:
	;
	m.G0 = v10 + int32(1488)
	return
L4:
	;
	F_RegisterBackgroundWorker(m, v10+int32(16))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+1480)) = v55
	v61 = F_RegisterDynamicBackgroundWorker(m, v10+int32(16), v10+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	goto L3
L9:
	;
	if v61 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	goto L12
L11:
	;
	if v137 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[10]))
	if v76 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v106
	v137 = int32(0)
	goto L11
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[11]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	v87 = F_LWLockAcquire(m, v83+int32(_a_F_apw_start_leader_worker_0), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	v92 = v80 + v81*int32(1488)
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)+24))
	if v89 == v93 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	F_LWLockRelease(m, v108+int32(_a_F_apw_start_leader_worker_0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L25
	}
L20:
	;
	v96 = v92 + int32(16)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	F_LWLockRelease(m, v100+int32(_a_F_apw_start_leader_worker_0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v137 = int32(2)
	goto L11
L25:
	;
	if v106 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L13
L27:
	;
	if v106 != 0 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[13]))
	v122 = F_WaitLatch(m, v118, int32(17), int32(0), int32(134217734))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	v137 = int32(2)
	goto L11
L31:
	;
	if v122&int32(16) != 0 {
		v137 = int32(3)
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[13]))
	v128 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v128
	v133 = base.AtomicRmwOr32(m, v128, int32(_a_F_apw_start_leader_worker_1), v128)
	goto L33
L33:
	;
	goto L12
L34:
	;
	goto L3
L35:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_apw_start_leader_worker_2), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errhint(m, int32(_a_F_apw_start_leader_worker_3), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_apw_start_leader_worker_4), int32(944), int32(_a_F_apw_start_leader_worker_5))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_apw_start_leader_worker_6), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errhint(m, int32(_a_F_apw_start_leader_worker_7), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_apw_start_leader_worker_4), int32(951), int32(_a_F_apw_start_leader_worker_5))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_arrayexpr_cleanup_fn(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	F_list_free(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_pfree(m, v2)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_assignOperTypes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v12 = F_SearchSysCache1(m, int32(40), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)))
			if v17 != int32(98) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_errcode(m, int32(117833860))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_assignOperTypes_0), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1174), int32(_a_F_assignOperTypes_2))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v20 != 0 {
					v22 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
						if v24 != 0 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v48 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
							} else {
							}
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v53 == int32(0) {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
							} else {
							}
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									v32 = F_get_am_name(m, l1)
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v32
										F_errmsg(m, int32(_a_F_assignOperTypes_3), v8+int32(16))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1192), int32(_a_F_assignOperTypes_2))
											mBase = m.M
											v44 = m.ExcPending
											if v44 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
					if v45 != int32(16) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_assignOperTypes_4), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1202), int32(_a_F_assignOperTypes_2))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v48 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+80))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
						} else {
						}
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v53 == int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+84))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v56
						} else {
						}
						F_ReleaseCatCache(m, v12)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
				F_errmsg_internal(m, int32(_a_F_assignOperTypes_5), v8)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1165), int32(_a_F_assignOperTypes_2))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_assign_debug_io_direct(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_debug_io_direct[0])) = v4
	return
}
func F_assign_maintenance_io_concurrency(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	*(*int32)(unsafe.Add(mBase, _c_F_assign_maintenance_io_concurrency[0])) = l0
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_assign_maintenance_io_concurrency[1]))
	if v6 == int32(13) {
		v9 = int32(_a_F_assign_maintenance_io_concurrency_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_assign_maintenance_io_concurrency[2]))
		*(*int32)(unsafe.Add(mBase, _c_F_assign_maintenance_io_concurrency[2])) = v11 + int32(1)
	} else {
	}
	return
}
func F_assign_session_replication_role(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_session_replication_role[0]))
	if l0 != v4 {
		v6 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_assign_session_replication_role[1]))
		if base.B2i32(v10 == v6)|base.B2i32(v10 == int32(_a_F_assign_session_replication_role_0)) == v6 {
			v18 = v10
			for {
				v22 = v18 - int32(5)
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v23 != int32(1) {
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(96))))
					if v28 != 0 {
						v29 = F_stmt_requires_parse_analysis(m, v28)
						mBase = m.M
						if v29 != 0 {
							v39 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v39)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(12))))
							if v43 == v39 {
							} else {
								v46 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v43)+10)) = uint8(v46)
							}
						} else {
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(92))))
						if v32 == int32(0) {
						} else {
							v35 = F_query_requires_rewrite_plan(m, v32)
							mBase = m.M
							if v35 == int32(0) {
							} else {
								v39 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v39)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(12))))
								if v43 == v39 {
								} else {
									v46 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v43)+10)) = uint8(v46)
								}
							}
						}
					}
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				if v50 != int32(_a_F_assign_session_replication_role_0) {
					v18 = v50
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v57 = *(*int32)(unsafe.Add(mBase, _c_F_assign_session_replication_role[2]))
		v58 = int32(0)
		if base.B2i32(v57 == v58)|base.B2i32(v57 == int32(_a_F_assign_session_replication_role_1)) == v58 {
			v65 = v57
			for {
				v70 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v65-int32(16)))) = uint8(v70)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
				if v72 != int32(_a_F_assign_session_replication_role_1) {
					v65 = v72
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
	} else {
	}
	return
}
func F_assign_syslog_ident(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_ident[0]))
	if v4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v7 == int32(0))|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_syslog_ident[1])))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	if v28-v29 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v13 = v4
	v14 = l0
	goto L8
L8:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v28 = v18
	v29 = v17
	goto L6
L10:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L4
L13:
	;
	v36 = m.G0
	v37 = int32(16)
	v38 = v36 - v37
	m.G0 = v38
	v40 = int32(_a_F_assign_syslog_ident_0)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_ident[2]))
	v42 = F_close(m, v41)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_ident[2])) = int32(-1)
	m.G0 = v38 + v37
	goto L16
L14:
	;
	v55 = v4
	goto L15
L15:
	;
	F_emscripten_builtin_free(m, v55)
	mBase = m.M
	v60 = F_strlen(m, l0)
	mBase = m.M
	v62 = v60 + int32(1)
	v63 = F_emscripten_builtin_malloc(m, v62)
	mBase = m.M
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_syslog_ident[1])) = uint8(v51)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_ident[0]))
	v55 = v54
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_ident[0])) = v68
	goto L1
L18:
	;
	v68 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v67 = F___memcpy(m, v63, l0, v62)
	mBase = m.M
	v68 = v67
	goto L17
}
func F_atan(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v23 float64
	_ = v23
	var v30 float64
	_ = v30
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v78 float64
	_ = v78
	var v95 float64
	_ = v95
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	v8 = base.I64_reinterpret_f64(l0)
	v13 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1141899264)) <= base.Ui32(v13) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
			v23 = l0
		} else {
			v23 = base.F64_copysign(float64(1.5707963267948966), l0)
		}
		return v23
	} else {
		if base.Ui32(v13) <= base.Ui32(int32(1071382527)) {
			if base.Ui32(int32(1044381696)) <= base.Ui32(v13) {
				v61 = l0
				v62 = int32(-1)
				v63 = base.F64_mul(v61, v61)
				v64 = base.F64_mul(v63, v63)
				v78 = base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v95 = base.F64_mul(v63, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v13) <= base.Ui32(int32(1071382527)) {
					return base.F64_sub(v61, base.F64_mul(v61, base.F64_add(v78, v95)))
				} else {
					v103 = v62 << (uint(int32(3)) % 32)
					v104 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_atan[0])))
					v107 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_atan[1])))
					v110 = base.F64_sub(v104, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v107), v61))
					if v8 < int64(0) {
						v114 = base.F64_neg(v110)
					} else {
						v114 = v110
					}
					v115 = v114
					return v115
				}
			} else {
				v115 = l0
				return v115
			}
		} else {
			v30 = base.F64_abs(l0)
			if base.Ui32(v13) <= base.Ui32(int32(1072889855)) {
				if base.Ui32(v13) <= base.Ui32(int32(1072037887)) {
					v61 = base.F64_div(base.F64_add(base.F64_add(v30, v30), float64(-1)), base.F64_add(v30, float64(2)))
					v62 = int32(0)
				} else {
					v61 = base.F64_div(base.F64_add(v30, float64(-1)), base.F64_add(v30, float64(1)))
					v62 = int32(1)
				}
			} else {
				if base.Ui32(v13) <= base.Ui32(int32(1073971199)) {
					v61 = base.F64_div(base.F64_add(v30, float64(-1.5)), base.F64_add(base.F64_mul(v30, float64(1.5)), float64(1)))
					v62 = int32(2)
				} else {
					v61 = base.F64_div(float64(-1), v30)
					v62 = int32(3)
				}
			}
			v63 = base.F64_mul(v61, v61)
			v64 = base.F64_mul(v63, v63)
			v78 = base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
			v95 = base.F64_mul(v63, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, base.F64_add(base.F64_mul(v64, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
			if base.Ui32(v13) <= base.Ui32(int32(1071382527)) {
				return base.F64_sub(v61, base.F64_mul(v61, base.F64_add(v78, v95)))
			} else {
				v103 = v62 << (uint(int32(3)) % 32)
				v104 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_atan[0])))
				v107 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_atan[1])))
				v110 = base.F64_sub(v104, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v107), v61))
				if v8 < int64(0) {
					v114 = base.F64_neg(v110)
				} else {
					v114 = v110
				}
				v115 = v114
				return v115
			}
		}
	}
}
func F_atan2(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v55 float64
	_ = v55
	var v61 float64
	_ = v61
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v109 float64
	_ = v109
	var v126 float64
	_ = v126
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v155 float64
	_ = v155
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v195 float64
	_ = v195
	var v213 float64
	_ = v213
	var v220 int64
	_ = v220
	var v225 int32
	_ = v225
	var v235 float64
	_ = v235
	var v241 float64
	_ = v241
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v289 float64
	_ = v289
	var v306 float64
	_ = v306
	var v313 int32
	_ = v313
	var v314 float64
	_ = v314
	var v317 float64
	_ = v317
	var v320 float64
	_ = v320
	var v324 float64
	_ = v324
	var v325 float64
	_ = v325
	var v335 float64
	_ = v335
	var v336 float64
	_ = v336
	var v353 float64
	_ = v353
	var v354 float64
	_ = v354
	v10 = int64(9223372036854775807)
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(l0)&v10) < base.Ui64(int64(9218868437227405313)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(l1)&v10) <= base.Ui64(int64(9218868437227405312))) == int32(0) {
		return base.F64_add(l0, l1)
	} else {
		v24 = base.I64_reinterpret_f64(l1)
		v27 = base.I32_wrap_i64(int64(base.Ui64(v24) >> (uint(int64(32)) % 64)))
		v30 = base.I32_wrap_i64(v24)
		if v27-int32(1072693248)|v30 == int32(0) {
			v40 = base.I64_reinterpret_f64(l0)
			v45 = base.I32_wrap_i64(int64(base.Ui64(v40)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1141899264)) <= base.Ui32(v45) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
					v55 = l0
				} else {
					v55 = base.F64_copysign(float64(1.5707963267948966), l0)
				}
				v155 = v55
			} else {
				if base.Ui32(v45) <= base.Ui32(int32(1071382527)) {
					if base.Ui32(int32(1044381696)) <= base.Ui32(v45) {
						v92 = l0
						v93 = int32(-1)
						v94 = base.F64_mul(v92, v92)
						v95 = base.F64_mul(v94, v94)
						v109 = base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
						v126 = base.F64_mul(v94, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
						if base.Ui32(v45) <= base.Ui32(int32(1071382527)) {
							v155 = base.F64_sub(v92, base.F64_mul(v92, base.F64_add(v109, v126)))
						} else {
							v133 = v93 << (uint(int32(3)) % 32)
							v134 = *(*float64)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_atan2[0])))
							v137 = *(*float64)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_atan2[1])))
							v140 = base.F64_sub(v134, base.F64_sub(base.F64_sub(base.F64_mul(v92, base.F64_add(v109, v126)), v137), v92))
							if v40 < int64(0) {
								v144 = base.F64_neg(v140)
							} else {
								v144 = v140
							}
							v145 = v144
							v155 = v145
						}
					} else {
						v145 = l0
						v155 = v145
					}
				} else {
					v61 = base.F64_abs(l0)
					if base.Ui32(v45) <= base.Ui32(int32(1072889855)) {
						if base.Ui32(v45) <= base.Ui32(int32(1072037887)) {
							v92 = base.F64_div(base.F64_add(base.F64_add(v61, v61), float64(-1)), base.F64_add(v61, float64(2)))
							v93 = int32(0)
						} else {
							v92 = base.F64_div(base.F64_add(v61, float64(-1)), base.F64_add(v61, float64(1)))
							v93 = int32(1)
						}
					} else {
						if base.Ui32(v45) <= base.Ui32(int32(1073971199)) {
							v92 = base.F64_div(base.F64_add(v61, float64(-1.5)), base.F64_add(base.F64_mul(v61, float64(1.5)), float64(1)))
							v93 = int32(2)
						} else {
							v92 = base.F64_div(float64(-1), v61)
							v93 = int32(3)
						}
					}
					v94 = base.F64_mul(v92, v92)
					v95 = base.F64_mul(v94, v94)
					v109 = base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
					v126 = base.F64_mul(v94, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
					if base.Ui32(v45) <= base.Ui32(int32(1071382527)) {
						v155 = base.F64_sub(v92, base.F64_mul(v92, base.F64_add(v109, v126)))
					} else {
						v133 = v93 << (uint(int32(3)) % 32)
						v134 = *(*float64)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_atan2[0])))
						v137 = *(*float64)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_atan2[1])))
						v140 = base.F64_sub(v134, base.F64_sub(base.F64_sub(base.F64_mul(v92, base.F64_add(v109, v126)), v137), v92))
						if v40 < int64(0) {
							v144 = base.F64_neg(v140)
						} else {
							v144 = v140
						}
						v145 = v144
						v155 = v145
					}
				}
			}
			return v155
		} else {
			v160 = int32(base.Ui32(v27)>>(uint(int32(30))%32)) & int32(2)
			v161 = base.I64_reinterpret_f64(l0)
			v165 = v160 | base.I32_wrap_i64(int64(base.Ui64(v161)>>(uint(int64(63))%64)))
			v170 = base.I32_wrap_i64(int64(base.Ui64(v161)>>(uint(int64(32))%64))) & int32(2147483647)
			if v170|base.I32_wrap_i64(v161) == int32(0) {
				switch v165 - int32(2) {
				case 0:
					return float64(3.141592653589793)
				case 1:
					return float64(-3.141592653589793)
				default:
					v354 = l0
					return v354
				}
			} else {
				v182 = v27 & int32(2147483647)
				if v182|v30 == int32(0) {
					return base.F64_copysign(float64(1.5707963267948966), l0)
				} else {
					if v182 == int32(2146435072) {
						if v170 != int32(2146435072) {
							v353 = *(*float64)(unsafe.Add(mBase, uint32(v165<<(uint(int32(3))%32))+uint32(_c_F_atan2[2])))
							v354 = v353
							return v354
						} else {
							v195 = *(*float64)(unsafe.Add(mBase, uint32(v165<<(uint(int32(3))%32))+uint32(_c_F_atan2[3])))
							return v195
						}
					} else {
						if base.B2i32(v170 != int32(2146435072))&base.B2i32(base.Ui32(v170) <= base.Ui32(v182+int32(67108864))) == int32(0) {
							return base.F64_copysign(float64(1.5707963267948966), l0)
						} else {
							if v160 != 0 {
								if base.Ui32(v170+int32(67108864)) < base.Ui32(v182) {
									v336 = float64(0)
								} else {
									v213 = base.F64_abs(base.F64_div(l0, l1))
									v220 = base.I64_reinterpret_f64(v213)
									v225 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(32))%64))) & int32(2147483647)
									if base.Ui32(int32(1141899264)) <= base.Ui32(v225) {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v213)&int64(9223372036854775807)) {
											v235 = v213
										} else {
											v235 = base.F64_copysign(float64(1.5707963267948966), v213)
										}
										v335 = v235
									} else {
										if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
											if base.Ui32(int32(1044381696)) <= base.Ui32(v225) {
												v272 = v213
												v273 = int32(-1)
												v274 = base.F64_mul(v272, v272)
												v275 = base.F64_mul(v274, v274)
												v289 = base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
												v306 = base.F64_mul(v274, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
												if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
													v335 = base.F64_sub(v272, base.F64_mul(v272, base.F64_add(v289, v306)))
												} else {
													v313 = v273 << (uint(int32(3)) % 32)
													v314 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[0])))
													v317 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[1])))
													v320 = base.F64_sub(v314, base.F64_sub(base.F64_sub(base.F64_mul(v272, base.F64_add(v289, v306)), v317), v272))
													if v220 < int64(0) {
														v324 = base.F64_neg(v320)
													} else {
														v324 = v320
													}
													v325 = v324
													v335 = v325
												}
											} else {
												v325 = v213
												v335 = v325
											}
										} else {
											v241 = base.F64_abs(v213)
											if base.Ui32(v225) <= base.Ui32(int32(1072889855)) {
												if base.Ui32(v225) <= base.Ui32(int32(1072037887)) {
													v272 = base.F64_div(base.F64_add(base.F64_add(v241, v241), float64(-1)), base.F64_add(v241, float64(2)))
													v273 = int32(0)
												} else {
													v272 = base.F64_div(base.F64_add(v241, float64(-1)), base.F64_add(v241, float64(1)))
													v273 = int32(1)
												}
											} else {
												if base.Ui32(v225) <= base.Ui32(int32(1073971199)) {
													v272 = base.F64_div(base.F64_add(v241, float64(-1.5)), base.F64_add(base.F64_mul(v241, float64(1.5)), float64(1)))
													v273 = int32(2)
												} else {
													v272 = base.F64_div(float64(-1), v241)
													v273 = int32(3)
												}
											}
											v274 = base.F64_mul(v272, v272)
											v275 = base.F64_mul(v274, v274)
											v289 = base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
											v306 = base.F64_mul(v274, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
											if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
												v335 = base.F64_sub(v272, base.F64_mul(v272, base.F64_add(v289, v306)))
											} else {
												v313 = v273 << (uint(int32(3)) % 32)
												v314 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[0])))
												v317 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[1])))
												v320 = base.F64_sub(v314, base.F64_sub(base.F64_sub(base.F64_mul(v272, base.F64_add(v289, v306)), v317), v272))
												if v220 < int64(0) {
													v324 = base.F64_neg(v320)
												} else {
													v324 = v320
												}
												v325 = v324
												v335 = v325
											}
										}
									}
									v336 = v335
								}
							} else {
								v213 = base.F64_abs(base.F64_div(l0, l1))
								v220 = base.I64_reinterpret_f64(v213)
								v225 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(32))%64))) & int32(2147483647)
								if base.Ui32(int32(1141899264)) <= base.Ui32(v225) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v213)&int64(9223372036854775807)) {
										v235 = v213
									} else {
										v235 = base.F64_copysign(float64(1.5707963267948966), v213)
									}
									v335 = v235
								} else {
									if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
										if base.Ui32(int32(1044381696)) <= base.Ui32(v225) {
											v272 = v213
											v273 = int32(-1)
											v274 = base.F64_mul(v272, v272)
											v275 = base.F64_mul(v274, v274)
											v289 = base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
											v306 = base.F64_mul(v274, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
											if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
												v335 = base.F64_sub(v272, base.F64_mul(v272, base.F64_add(v289, v306)))
											} else {
												v313 = v273 << (uint(int32(3)) % 32)
												v314 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[0])))
												v317 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[1])))
												v320 = base.F64_sub(v314, base.F64_sub(base.F64_sub(base.F64_mul(v272, base.F64_add(v289, v306)), v317), v272))
												if v220 < int64(0) {
													v324 = base.F64_neg(v320)
												} else {
													v324 = v320
												}
												v325 = v324
												v335 = v325
											}
										} else {
											v325 = v213
											v335 = v325
										}
									} else {
										v241 = base.F64_abs(v213)
										if base.Ui32(v225) <= base.Ui32(int32(1072889855)) {
											if base.Ui32(v225) <= base.Ui32(int32(1072037887)) {
												v272 = base.F64_div(base.F64_add(base.F64_add(v241, v241), float64(-1)), base.F64_add(v241, float64(2)))
												v273 = int32(0)
											} else {
												v272 = base.F64_div(base.F64_add(v241, float64(-1)), base.F64_add(v241, float64(1)))
												v273 = int32(1)
											}
										} else {
											if base.Ui32(v225) <= base.Ui32(int32(1073971199)) {
												v272 = base.F64_div(base.F64_add(v241, float64(-1.5)), base.F64_add(base.F64_mul(v241, float64(1.5)), float64(1)))
												v273 = int32(2)
											} else {
												v272 = base.F64_div(float64(-1), v241)
												v273 = int32(3)
											}
										}
										v274 = base.F64_mul(v272, v272)
										v275 = base.F64_mul(v274, v274)
										v289 = base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
										v306 = base.F64_mul(v274, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, base.F64_add(base.F64_mul(v275, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
										if base.Ui32(v225) <= base.Ui32(int32(1071382527)) {
											v335 = base.F64_sub(v272, base.F64_mul(v272, base.F64_add(v289, v306)))
										} else {
											v313 = v273 << (uint(int32(3)) % 32)
											v314 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[0])))
											v317 = *(*float64)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_atan2[1])))
											v320 = base.F64_sub(v314, base.F64_sub(base.F64_sub(base.F64_mul(v272, base.F64_add(v289, v306)), v317), v272))
											if v220 < int64(0) {
												v324 = base.F64_neg(v320)
											} else {
												v324 = v320
											}
											v325 = v324
											v335 = v325
										}
									}
								}
								v336 = v335
							}
							switch v165 - int32(1) {
							case 0:
								return base.F64_neg(v336)
							case 1:
								return base.F64_sub(float64(3.141592653589793), base.F64_add(v336, float64(-1.2246467991473532e-16)))
							case 2:
								return base.F64_add(base.F64_add(v336, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
							default:
								v354 = v336
								return v354
							}
						}
					}
				}
			}
		}
	}
}
