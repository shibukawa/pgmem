package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		v18 = base.B2i32(base.Ui32(v6) <= base.Ui32(l0))
	} else {
		v18 = base.B2i32(int32(0) <= l0-v6)
	}
	if v18 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
		v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+12)))
		v26 = int32(3)
		v28 = l0 + int32(1)
		if base.Ui32(v28) <= base.Ui32(v26) {
			v31 = v26
		} else {
			v31 = v28
		}
		if base.Ui32(v31) < base.Ui32(v6) {
			v33 = (v21 + int64(1)) & int64(4294967295)
		} else {
			v33 = v21
		}
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
		v39 = F_LWLockAcquire(m, v35+int32(384), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[0]))
			*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = base.I64_extend_i32_u(v31) | v33<<(uint(int64(32))%64)
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceNextFullTransactionIdPastXid[1]))
			F_LWLockRelease(m, v49+int32(384))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				return
			}
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	v13 = v11 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if l2 != v14 {
		v16 = F_superuser(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(65536)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v60 = F_heap_getattr_7(m, l1, int32(7), v57, v7+int32(-49))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v62 == int32(0) {
						v65 = F_pg_detoast_datum(m, v60)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
							v68 = F_aclnewowner(m, v65, v67, l2)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v68
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v71)
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										F_changeDependencyOnOwner(m, int32(1417), v87, l2)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
											if v93 != 0 {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v96 = int32(0)
												F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													m.G0 = v9 - int32(-64)
													return
												}
											} else {
												m.G0 = v9 - int32(-64)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								F_changeDependencyOnOwner(m, int32(1417), v87, l2)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
									if v93 != 0 {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v96 = int32(0)
										F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											m.G0 = v9 - int32(-64)
											return
										}
									} else {
										m.G0 = v9 - int32(-64)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
				v22 = F_object_ownercheck(m, int32(1417), v19, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if v22 == int32(0) {
						F_aclcheck_error(m, int32(2), int32(17), v13+int32(4))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
							F_check_can_set_role(m, v33, l2)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
								v39 = F_object_aclcheck(m, int32(2328), v37, l2, int64(256))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									if v39 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(65536)
										*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v60 = F_heap_getattr_7(m, l1, int32(7), v57, v7+int32(-49))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											if v62 == int32(0) {
												v65 = F_pg_detoast_datum(m, v60)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
													v68 = F_aclnewowner(m, v65, v67, l2)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v68
														v71 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v71)
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
														v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(1417), v87, l2)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																	if v93 != 0 {
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v96 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return
																		} else {
																			m.G0 = v9 - int32(-64)
																			return
																		}
																	} else {
																		m.G0 = v9 - int32(-64)
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														F_changeDependencyOnOwner(m, int32(1417), v87, l2)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
															if v93 != 0 {
																v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v96 = int32(0)
																F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	m.G0 = v9 - int32(-64)
																	return
																}
															} else {
																m.G0 = v9 - int32(-64)
																return
															}
														}
													}
												}
											}
										}
									} else {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
										v45 = F_GetForeignDataWrapper(m, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
											F_aclcheck_error(m, v39, int32(16), v47)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(65536)
												*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
												v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
												v60 = F_heap_getattr_7(m, l1, int32(7), v57, v7+int32(-49))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
													if v62 == int32(0) {
														v65 = F_pg_detoast_datum(m, v60)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
															v68 = F_aclnewowner(m, v65, v67, l2)
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v68
																v71 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v71)
																v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
																v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return
																} else {
																	F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return
																	} else {
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		F_changeDependencyOnOwner(m, int32(1417), v87, l2)
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return
																		} else {
																			v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																			if v93 != 0 {
																				v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v96 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																				mBase = m.M
																				v100 = m.ExcPending
																				if v100 != 0 {
																					return
																				} else {
																					m.G0 = v9 - int32(-64)
																					return
																				}
																			} else {
																				m.G0 = v9 - int32(-64)
																				return
																			}
																		}
																	}
																}
															}
														}
													} else {
														v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
														v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return
														} else {
															F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return
															} else {
																v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(1417), v87, l2)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return
																} else {
																	v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																	if v93 != 0 {
																		v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v96 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return
																		} else {
																			m.G0 = v9 - int32(-64)
																			return
																		}
																	} else {
																		m.G0 = v9 - int32(-64)
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
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[1]))
						F_check_can_set_role(m, v33, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
							v39 = F_object_aclcheck(m, int32(2328), v37, l2, int64(256))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								if v39 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(65536)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v60 = F_heap_getattr_7(m, l1, int32(7), v57, v7+int32(-49))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										if v62 == int32(0) {
											v65 = F_pg_detoast_datum(m, v60)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
												v68 = F_aclnewowner(m, v65, v67, l2)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v68
													v71 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v71)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(1417), v87, l2)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																if v93 != 0 {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v96 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		m.G0 = v9 - int32(-64)
																		return
																	}
																} else {
																	m.G0 = v9 - int32(-64)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													F_changeDependencyOnOwner(m, int32(1417), v87, l2)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
														if v93 != 0 {
															v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v96 = int32(0)
															F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																m.G0 = v9 - int32(-64)
																return
															}
														} else {
															m.G0 = v9 - int32(-64)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
									v45 = F_GetForeignDataWrapper(m, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
										F_aclcheck_error(m, v39, int32(16), v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(65536)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l2
											v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v60 = F_heap_getattr_7(m, l1, int32(7), v57, v7+int32(-49))
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
												if v62 == int32(0) {
													v65 = F_pg_detoast_datum(m, v60)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
														v68 = F_aclnewowner(m, v65, v67, l2)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v68
															v71 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v71)
															v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
															v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return
															} else {
																F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return
																} else {
																	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(1417), v87, l2)
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return
																	} else {
																		v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																		if v93 != 0 {
																			v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v96 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																			mBase = m.M
																			v100 = m.ExcPending
																			if v100 != 0 {
																				return
																			} else {
																				m.G0 = v9 - int32(-64)
																				return
																			}
																		} else {
																			m.G0 = v9 - int32(-64)
																			return
																		}
																	}
																}
															}
														}
													}
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v80 = F_heap_modify_tuple(m, l1, v73, v7+int32(-32), v7+int32(-40), v7+int32(-48))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_CatalogTupleUpdate(m, l0, v80+int32(4), v80)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(1417), v87, l2)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
																if v93 != 0 {
																	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v96 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		m.G0 = v9 - int32(-64)
																		return
																	}
																} else {
																	m.G0 = v9 - int32(-64)
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
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_AlterForeignServerOwner_internal[0]))
		if v93 != 0 {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v96 = int32(0)
			F_RunObjectPostAlterHook(m, int32(1417), v95, v96, v96, v96)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				m.G0 = v9 - int32(-64)
				return
			}
		} else {
			m.G0 = v9 - int32(-64)
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
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v384 = int32(2)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v379-int32(1023)) < base.Ui32(v384) {
		goto L66
	} else {
		goto L67
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v11
	goto L1
L3:
	;
	goto L4
L4:
	;
	v12 = int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v14-int32(1021)) < base.Ui32(v12) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
	v379 = v369
	goto L1
L6:
	;
	v23 = v14
	v25 = v12
	v27 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v13))) = v8
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v369 = v364 + int32(4)
	goto L5
L9:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v23) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v369 = v360
	goto L5
L11:
	;
	v30 = int32(1024)
	v37 = int32(-1636607408)
	goto L16
L12:
	;
	v352 = v23
	goto L13
L13:
	;
	v354 = int32(1024) - v352
	if base.Ui32(v25) < base.Ui32(v354) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v337^v329-base.I32_rotl(v337, int32(24))))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L57
	} else {
		goto L58
	}
L15:
	;
	if v13&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	goto L15
L19:
	;
	v315 = int32(14)
	v317 = v311 ^ v312 - base.I32_rotl(v311, v315)
	v321 = v317 ^ v310 - base.I32_rotl(v317, int32(11))
	v325 = v321 ^ v311 - base.I32_rotl(v321, int32(25))
	v329 = v325 ^ v317 - base.I32_rotl(v325, int32(16))
	v333 = v329 ^ v321 - base.I32_rotl(v329, int32(4))
	v337 = v333 ^ v325 - base.I32_rotl(v333, v315)
	goto L14
L20:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v310 = v302 + v305
	v311 = v303
	v312 = v304
	goto L19
L21:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	v302 = v298<<(uint(int32(8))%32) + v295
	v303 = v296
	v304 = v297
	goto L20
L22:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	v295 = v291<<(uint(int32(16))%32) + v288
	v296 = v289
	v297 = v290
	goto L21
L23:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
	v288 = v284<<(uint(int32(24))%32) + v120
	v289 = v282
	v290 = v283
	goto L22
L24:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+4)))
	v282 = v278 + v280
	v283 = v279
	goto L23
L25:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+5)))
	v278 = v274<<(uint(int32(8))%32) + v272
	v279 = v273
	goto L24
L26:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
	v272 = v268<<(uint(int32(16))%32) + v266
	v273 = v267
	goto L25
L27:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+7)))
	v266 = v262<<(uint(int32(24))%32) + v121
	v267 = v261
	goto L26
L28:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+8)))
	v261 = v257<<(uint(int32(8))%32) + v256
	goto L27
L29:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+9)))
	v256 = v252<<(uint(int32(16))%32) + v251
	goto L28
L30:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+10)))
	v251 = v247<<(uint(int32(24))%32) + v125
	goto L29
L31:
	;
	goto L34
L32:
	;
	goto L33
L33:
	;
	goto L40
L34:
	;
	v83 = v13
	v84 = v30
	v86 = v37
	v87 = v37
	v88 = v37
	goto L37
L36:
	;
	switch v129 - int32(1) {
	case 0:
		v302 = v120
		v303 = v121
		v304 = v125
		goto L20
	case 1:
		v295 = v120
		v296 = v121
		v297 = v125
		goto L21
	case 2:
		v288 = v120
		v289 = v121
		v290 = v125
		goto L22
	case 3:
		v282 = v121
		v283 = v125
		goto L23
	case 4:
		v278 = v121
		v279 = v125
		goto L24
	case 5:
		v272 = v121
		v273 = v125
		goto L25
	case 6:
		v266 = v121
		v267 = v125
		goto L26
	case 7:
		v261 = v125
		goto L27
	case 8:
		v256 = v125
		goto L28
	case 9:
		v251 = v125
		goto L29
	case 10:
		goto L30
	default:
		v310 = v120
		v311 = v121
		v312 = v125
		goto L19
	}
L37:
	;
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
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v143 = v13
	v144 = v30
	v146 = v37
	v147 = v37
	v148 = v37
	goto L43
L42:
	;
	switch v189 - int32(1) {
	case 0:
		v244 = v180
		goto L46
	case 1:
		v239 = v180
		goto L47
	case 2:
		goto L48
	case 3:
		v232 = v181
		goto L49
	case 4:
		v229 = v181
		goto L50
	case 5:
		v224 = v181
		goto L51
	case 6:
		goto L52
	case 7:
		v215 = v185
		goto L53
	case 8:
		v210 = v185
		goto L54
	case 9:
		v205 = v185
		goto L55
	case 10:
		goto L56
	default:
		v310 = v180
		v311 = v181
		v312 = v185
		goto L19
	}
L43:
	;
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
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v310 = v244 + v245
	v311 = v181
	v312 = v185
	goto L19
L47:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v244 = v240<<(uint(int32(8))%32) + v239
	goto L46
L48:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	v239 = v235<<(uint(int32(16))%32) + v180
	goto L47
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v310 = v233 + v180
	v311 = v232
	v312 = v185
	goto L19
L50:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
	v232 = v229 + v230
	goto L49
L51:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
	v229 = v225<<(uint(int32(8))%32) + v224
	goto L50
L52:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	v224 = v220<<(uint(int32(16))%32) + v181
	goto L51
L53:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v310 = v216 + v180
	v311 = v218 + v181
	v312 = v215
	goto L19
L54:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	v215 = v211<<(uint(int32(8))%32) + v210
	goto L53
L55:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	v210 = v206<<(uint(int32(16))%32) + v205
	goto L54
L56:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
	v205 = v201<<(uint(int32(24))%32) + v185
	goto L55
L57:
	;
	return
L58:
	;
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v349
	v352 = int32(8)
	goto L13
L59:
	;
	v356 = v25
	goto L61
L60:
	;
	v356 = v354
	goto L61
L61:
	;
	if v356 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v352+v13, v27, v356)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v360 = v352 + v356
	v361 = v25 - v356
	if v361 != 0 {
		v23 = v360
		v25 = v361
		v27 = v356 + v27
		goto L9
	} else {
		goto L65
	}
L65:
	;
	goto L10
L66:
	;
	v391 = l1
	v392 = v379
	v394 = v384
	goto L69
L67:
	;
	goto L68
L68:
	;
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*uint16)(unsafe.Add(mBase, uint32(v379+v385))) = uint16(v733)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v735 + int32(2)
	return
L69:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v392) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729
	return
L71:
	;
	v399 = int32(1024)
	v406 = int32(-1636607408)
	goto L76
L72:
	;
	v721 = v392
	goto L73
L73:
	;
	v723 = int32(1024) - v721
	if base.Ui32(v394) < base.Ui32(v723) {
		goto L118
	} else {
		goto L119
	}
L74:
	;
	v716 = F_Int64GetDatum(m, base.I64_extend_i32_u(v706)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v706^v698-base.I32_rotl(v706, int32(24))))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L57
	} else {
		goto L117
	}
L75:
	;
	if v385&int32(3) != 0 {
		goto L91
	} else {
		goto L92
	}
L76:
	;
	goto L75
L79:
	;
	v684 = int32(14)
	v686 = v680 ^ v681 - base.I32_rotl(v680, v684)
	v690 = v686 ^ v679 - base.I32_rotl(v686, int32(11))
	v694 = v690 ^ v680 - base.I32_rotl(v690, int32(25))
	v698 = v694 ^ v686 - base.I32_rotl(v694, int32(16))
	v702 = v698 ^ v690 - base.I32_rotl(v698, int32(4))
	v706 = v702 ^ v694 - base.I32_rotl(v702, v684)
	goto L74
L80:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	v679 = v671 + v674
	v680 = v672
	v681 = v673
	goto L79
L81:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+1)))
	v671 = v667<<(uint(int32(8))%32) + v664
	v672 = v665
	v673 = v666
	goto L80
L82:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+2)))
	v664 = v660<<(uint(int32(16))%32) + v657
	v665 = v658
	v666 = v659
	goto L81
L83:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+3)))
	v657 = v653<<(uint(int32(24))%32) + v489
	v658 = v651
	v659 = v652
	goto L82
L84:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+4)))
	v651 = v647 + v649
	v652 = v648
	goto L83
L85:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+5)))
	v647 = v643<<(uint(int32(8))%32) + v641
	v648 = v642
	goto L84
L86:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+6)))
	v641 = v637<<(uint(int32(16))%32) + v635
	v642 = v636
	goto L85
L87:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+7)))
	v635 = v631<<(uint(int32(24))%32) + v490
	v636 = v630
	goto L86
L88:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+8)))
	v630 = v626<<(uint(int32(8))%32) + v625
	goto L87
L89:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+9)))
	v625 = v621<<(uint(int32(16))%32) + v620
	goto L88
L90:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+10)))
	v620 = v616<<(uint(int32(24))%32) + v494
	goto L89
L91:
	;
	goto L94
L92:
	;
	goto L93
L93:
	;
	goto L100
L94:
	;
	v452 = v385
	v453 = v399
	v455 = v406
	v456 = v406
	v457 = v406
	goto L97
L96:
	;
	switch v498 - int32(1) {
	case 0:
		v671 = v489
		v672 = v490
		v673 = v494
		goto L80
	case 1:
		v664 = v489
		v665 = v490
		v666 = v494
		goto L81
	case 2:
		v657 = v489
		v658 = v490
		v659 = v494
		goto L82
	case 3:
		v651 = v490
		v652 = v494
		goto L83
	case 4:
		v647 = v490
		v648 = v494
		goto L84
	case 5:
		v641 = v490
		v642 = v494
		goto L85
	case 6:
		v635 = v490
		v636 = v494
		goto L86
	case 7:
		v630 = v494
		goto L87
	case 8:
		v625 = v494
		goto L88
	case 9:
		v620 = v494
		goto L89
	case 10:
		goto L90
	default:
		v679 = v489
		v680 = v490
		v681 = v494
		goto L79
	}
L97:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v460 = v459 + v456
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v464 = v463 + v457
	v466 = int32(4)
	v468 = v461 + v455 - v464 ^ base.I32_rotl(v464, v466)
	v472 = v460 - v468 ^ base.I32_rotl(v468, int32(6))
	v473 = v464 + v460
	v474 = v468 + v473
	v475 = v472 + v474
	v479 = v473 - v472 ^ base.I32_rotl(v472, int32(8))
	v483 = v474 - v479 ^ base.I32_rotl(v479, int32(16))
	v487 = v475 - v483 ^ base.I32_rotl(v483, int32(19))
	v488 = v479 + v475
	v489 = v483 + v488
	v490 = v487 + v489
	v494 = v488 - v487 ^ base.I32_rotl(v487, v466)
	v495 = int32(12)
	v496 = v452 + v495
	v498 = v453 - v495
	if base.Ui32(int32(11)) < base.Ui32(v498) {
		v452 = v496
		v453 = v498
		v455 = v489
		v456 = v490
		v457 = v494
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	goto L98
L100:
	;
	v512 = v385
	v513 = v399
	v515 = v406
	v516 = v406
	v517 = v406
	goto L103
L102:
	;
	switch v558 - int32(1) {
	case 0:
		v613 = v549
		goto L106
	case 1:
		v608 = v549
		goto L107
	case 2:
		goto L108
	case 3:
		v601 = v550
		goto L109
	case 4:
		v598 = v550
		goto L110
	case 5:
		v593 = v550
		goto L111
	case 6:
		goto L112
	case 7:
		v584 = v554
		goto L113
	case 8:
		v579 = v554
		goto L114
	case 9:
		v574 = v554
		goto L115
	case 10:
		goto L116
	default:
		v679 = v549
		v680 = v550
		v681 = v554
		goto L79
	}
L103:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v520 = v519 + v516
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v524 = v523 + v517
	v526 = int32(4)
	v528 = v521 + v515 - v524 ^ base.I32_rotl(v524, v526)
	v532 = v520 - v528 ^ base.I32_rotl(v528, int32(6))
	v533 = v524 + v520
	v534 = v528 + v533
	v535 = v532 + v534
	v539 = v533 - v532 ^ base.I32_rotl(v532, int32(8))
	v543 = v534 - v539 ^ base.I32_rotl(v539, int32(16))
	v547 = v535 - v543 ^ base.I32_rotl(v543, int32(19))
	v548 = v539 + v535
	v549 = v543 + v548
	v550 = v547 + v549
	v554 = v548 - v547 ^ base.I32_rotl(v547, v526)
	v555 = int32(12)
	v556 = v512 + v555
	v558 = v513 - v555
	if base.Ui32(int32(11)) < base.Ui32(v558) {
		v512 = v556
		v513 = v558
		v515 = v549
		v516 = v550
		v517 = v554
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	goto L104
L106:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v679 = v613 + v614
	v680 = v550
	v681 = v554
	goto L79
L107:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	v613 = v609<<(uint(int32(8))%32) + v608
	goto L106
L108:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+2)))
	v608 = v604<<(uint(int32(16))%32) + v549
	goto L107
L109:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v679 = v602 + v549
	v680 = v601
	v681 = v554
	goto L79
L110:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
	v601 = v598 + v599
	goto L109
L111:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+5)))
	v598 = v594<<(uint(int32(8))%32) + v593
	goto L110
L112:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+6)))
	v593 = v589<<(uint(int32(16))%32) + v550
	goto L111
L113:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	v679 = v585 + v549
	v680 = v587 + v550
	v681 = v584
	goto L79
L114:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+8)))
	v584 = v580<<(uint(int32(8))%32) + v579
	goto L113
L115:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+9)))
	v579 = v575<<(uint(int32(16))%32) + v574
	goto L114
L116:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+10)))
	v574 = v570<<(uint(int32(24))%32) + v554
	goto L115
L117:
	;
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v716)))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v718
	v721 = int32(8)
	goto L73
L118:
	;
	v725 = v394
	goto L120
L119:
	;
	v725 = v723
	goto L120
L120:
	;
	if v725 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	base.MemoryCopy(m, v721+v385, v391, v725)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v729 = v721 + v725
	v730 = v394 - v725
	if v730 != 0 {
		v391 = v391 + v725
		v392 = v729
		v394 = v730
		goto L69
	} else {
		goto L124
	}
L124:
	;
	goto L70
}
func F_ApplySetting(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v13 = v10 + int32(16)
	F_ScanKeyInit(m, v13, int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v10-int32(-64), int32(2), int32(3), int32(184), l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_systable_beginscan(m, l3, int32(2965), int32(1), l0, int32(2), v13)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = F_systable_getnext(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = v31
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_systable_endscan(m, v29)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v43 = F_heap_getattr_4(m, v34, v40, v10+int32(15))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = F_pg_detoast_datum(m, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v54 = F_systable_getnext(m, v29)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	F_ProcessGUCArray(m, v48, int32(5), l4, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v54 != 0 {
		v34 = v54
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	m.G0 = v10 + int32(112)
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v418 int64
	_ = v418
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int64
	_ = v893
	var v894 int64
	_ = v894
	var v904 int32
	_ = v904
	var v907 int64
	_ = v907
	var v914 int64
	_ = v914
	var v918 int64
	_ = v918
	var v919 int64
	_ = v919
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int64
	_ = v939
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int64
	_ = v962
	var v963 int64
	_ = v963
	var v973 int32
	_ = v973
	var v976 int64
	_ = v976
	var v983 int64
	_ = v983
	var v987 int64
	_ = v987
	var v988 int64
	_ = v988
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1015 int64
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int64
	_ = v1018
	var v1021 int64
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int64
	_ = v1030
	var v1033 int64
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int64
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1205 int32
	_ = v1205
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int64
	_ = v1233
	var v1234 int64
	_ = v1234
	var v1242 int64
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int64
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1296 int64
	_ = v1296
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1373 int32
	_ = v1373
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1438 int64
	_ = v1438
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1460 int32
	_ = v1460
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1507 int32
	_ = v1507
	var v1520 int32
	_ = v1520
	var v1531 int32
	_ = v1531
	var v1544 int32
	_ = v1544
	var v1554 int32
	_ = v1554
	var v1555 int64
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
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
	v1554 = int32(m.ExcTag)
	v1555 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1554 == int32(0) {
		goto L369
	} else {
		goto L370
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
	v553 = v17
	goto L9
L9:
	;
	if v553 != 0 {
		goto L146
	} else {
		goto L147
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
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[1])) = int32(3)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[0])) = int32(0)
	goto L12
L14:
	;
	v44 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L18
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[1]))
	v41 = F_GetBackendTypeDesc(m, v40)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_errmsg_internal(m, int32(_a_F_AutoVacLauncherMain_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[2]))
	if v56 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(385), int32(_a_F_AutoVacLauncherMain_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_pg_usleep(m, v56*int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v62 = int32(914)
	v64 = m.G0
	v66 = v64 - int32(32)
	m.G0 = v66
	switch int32(916) {
	case 0, 2:
		v76 = v62
		goto L29
	default:
		goto L30
	}
L27:
	;
	goto L26
L28:
	;
	v108 = int32(915)
	v110 = m.G0
	v112 = v110 - int32(32)
	m.G0 = v112
	switch int32(917) {
	case 0, 2:
		v122 = v108
		goto L42
	default:
		goto L43
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v76
	F_sigemptyset(m, v66+int32(16))
	mBase = m.M
	goto L32
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[3])) = v62
	v76 = int32(_a_F_AutoVacLauncherMain_4)
	goto L29
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(268435456)
	v88 = v66 + int32(12)
	goto L36
L34:
	;
	m.G0 = v66 + int32(32)
	goto L28
L36:
	;
	goto L37
L37:
	;
	if v88 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v94 = int32(20)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[4])) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v88)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[5])) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[6])) = v100
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L34
L41:
	;
	v154 = int32(916)
	v156 = m.G0
	v158 = v156 - int32(32)
	m.G0 = v158
	switch int32(918) {
	case 0, 2:
		v168 = v154
		goto L55
	default:
		goto L56
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v122
	F_sigemptyset(m, v112+int32(16))
	mBase = m.M
	goto L45
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[7])) = v108
	v122 = int32(_a_F_AutoVacLauncherMain_4)
	goto L42
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = int32(268435456)
	v134 = v112 + int32(12)
	goto L49
L47:
	;
	m.G0 = v112 + int32(32)
	goto L41
L49:
	;
	goto L50
L50:
	;
	if v134 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = int32(40)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[8])) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v134)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[9])) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[10])) = v146
	goto L53
L52:
	;
	goto L53
L53:
	;
	goto L47
L54:
	;
	v199 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[11])) = v199
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[12])) = v199
	v209 = v199
	goto L68
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v168
	F_sigemptyset(m, v158+int32(16))
	mBase = m.M
	goto L58
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[13])) = v154
	v168 = int32(_a_F_AutoVacLauncherMain_4)
	goto L55
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = int32(268435456)
	v180 = v158 + int32(12)
	goto L62
L60:
	;
	m.G0 = v158 + int32(32)
	goto L54
L62:
	;
	goto L63
L63:
	;
	if v180 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v187 = int32(300)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[14])) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[15])) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[16])) = v192
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L60
L67:
	;
	v288 = int32(-2)
	v290 = m.G0
	v292 = v290 - int32(32)
	m.G0 = v292
	switch int32(0) {
	case 0, 2:
		v302 = v288
		goto L74
	default:
		goto L75
	}
L68:
	;
	v211 = int32(40)
	v212 = v209 * v211
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[17]))) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[18]))) = v209
	v216 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[19]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[20]))) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[21]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[22]))) = v213
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+uint32(_c_F_AutoVacLauncherMain[23]))) = uint8(v213)
	v227 = v209 | int32(1)
	v229 = v227 * v211
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[17]))) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[18]))) = v227
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[19]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[20]))) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[21]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[22]))) = v213
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_AutoVacLauncherMain[23]))) = uint8(v213)
	v244 = v209 | int32(2)
	v246 = v244 * v211
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[17]))) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[18]))) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[19]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[20]))) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[21]))) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[22]))) = v213
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_AutoVacLauncherMain[23]))) = uint8(v213)
	if v209 != int32(20) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[24])) = uint8(v282)
	F_pqsignal_be(m, int32(14), int32(1769))
	mBase = m.M
	goto L67
L70:
	;
	v263 = v209 | int32(3)
	v265 = v263 * int32(40)
	v266 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[17]))) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[18]))) = v263
	v269 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[19]))) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[20]))) = v266
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[21]))) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[22]))) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_AutoVacLauncherMain[23]))) = uint8(v266)
	v209 = v209 + int32(4)
	goto L68
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v334 = int32(917)
	v336 = m.G0
	v338 = v336 - int32(32)
	m.G0 = v338
	switch int32(919) {
	case 0, 2:
		v348 = v334
		goto L87
	default:
		goto L88
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+12)) = v302
	F_sigemptyset(m, v292+int32(16))
	mBase = m.M
	goto L77
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[25])) = v288
	v302 = int32(_a_F_AutoVacLauncherMain_4)
	goto L74
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+24)) = int32(268435456)
	v314 = v292 + int32(12)
	goto L81
L79:
	;
	m.G0 = v292 + int32(32)
	goto L73
L81:
	;
	goto L82
L82:
	;
	if v314 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v321 = int32(260)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v314)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[26])) = v322
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v314)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[27])) = v324
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v314)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[28])) = v326
	goto L85
L84:
	;
	goto L85
L85:
	;
	goto L79
L86:
	;
	v380 = int32(918)
	v382 = m.G0
	v384 = v382 - int32(32)
	m.G0 = v384
	switch int32(920) {
	case 0, 2:
		v394 = v380
		goto L100
	default:
		goto L101
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+12)) = v348
	F_sigemptyset(m, v338+int32(16))
	mBase = m.M
	goto L90
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[29])) = v334
	v348 = int32(_a_F_AutoVacLauncherMain_4)
	goto L87
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+24)) = int32(268435456)
	v360 = v338 + int32(12)
	goto L94
L92:
	;
	m.G0 = v338 + int32(32)
	goto L86
L94:
	;
	goto L95
L95:
	;
	if v360 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v367 = int32(200)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v360)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[30])) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v360)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[31])) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v360)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[32])) = v372
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L92
L99:
	;
	v426 = int32(919)
	v428 = m.G0
	v430 = v428 - int32(32)
	m.G0 = v430
	switch int32(921) {
	case 0, 2:
		v440 = v426
		goto L113
	default:
		goto L114
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+12)) = v394
	F_sigemptyset(m, v384+int32(16))
	mBase = m.M
	goto L103
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[33])) = v380
	v394 = int32(_a_F_AutoVacLauncherMain_4)
	goto L100
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384)+24)) = int32(268435456)
	v406 = v384 + int32(12)
	goto L107
L105:
	;
	m.G0 = v384 + int32(32)
	goto L99
L107:
	;
	goto L108
L108:
	;
	if v406 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v413 = int32(240)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v406)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[34])) = v414
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v406)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[35])) = v416
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v406)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[36])) = v418
	goto L111
L110:
	;
	goto L111
L111:
	;
	goto L105
L112:
	;
	v472 = int32(0)
	v474 = m.G0
	v476 = v474 - int32(32)
	m.G0 = v476
	switch int32(2) {
	case 0, 2:
		v486 = v472
		goto L126
	default:
		goto L127
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+12)) = v440
	F_sigemptyset(m, v430+int32(16))
	mBase = m.M
	goto L116
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[37])) = v426
	v440 = int32(_a_F_AutoVacLauncherMain_4)
	goto L113
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = int32(268435456)
	v452 = v430 + int32(12)
	goto L120
L118:
	;
	m.G0 = v430 + int32(32)
	goto L112
L120:
	;
	goto L121
L121:
	;
	if v452 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v459 = int32(160)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v452)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[38])) = v460
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v452)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[39])) = v462
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v452)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[40])) = v464
	goto L124
L123:
	;
	goto L124
L124:
	;
	goto L118
L125:
	;
	F_InitProcess(m)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L138
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+12)) = v486
	F_sigemptyset(m, v476+int32(16))
	mBase = m.M
	goto L128
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[41])) = v472
	v486 = int32(_a_F_AutoVacLauncherMain_4)
	goto L126
L128:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+24)) = int32(268435457)
	v498 = v476 + int32(12)
	goto L133
L131:
	;
	m.G0 = v476 + int32(32)
	goto L125
L133:
	;
	goto L134
L134:
	;
	if v498 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v505 = int32(340)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v498)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[42])) = v506
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v498)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[43])) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v498)))
	*(*int64)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[44])) = v510
	goto L137
L136:
	;
	goto L137
L137:
	;
	goto L131
L138:
	;
	F_BaseInit(m)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	v521 = int32(0)
	F_InitPostgres(m, v521, v521, v521, v521, v521, v521)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[45])) = int32(2)
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[46]))
	v538 = F_AllocSetContextCreateInternal(m, v533, int32(_a_F_AutoVacLauncherMain_5), int32(0), int32(_a_F_AutoVacLauncherMain_6), int32(_a_F_AutoVacLauncherMain_7))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[47])) = v538
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[48])) = v538
	goto L142
L142:
	;
	v546 = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v546)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v546))) = v18 + int32(28)
	goto L145
L143:
	;
	v553 = int32(0)
	goto L9
L145:
	;
	goto L143
L146:
	;
	v554 = int32(_a_F_AutoVacLauncherMain_8)
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[49])) = v556 + int32(1)
	v561 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[50])) = v561
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[11])) = v561
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[12])) = v561
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[17])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[23])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[51])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[52])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[53])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[54])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[55])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[56])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[57])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[58])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[59])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[60])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[61])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[62])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[63])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[64])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[65])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[66])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[67])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[68])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[69])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[70])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[71])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[72])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[73])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[74])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[75])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[76])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[77])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[78])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[79])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[80])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[81])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[82])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[83])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[84])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[85])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[86])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[87])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[88])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[89])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[90])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[91])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[92])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[93])) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[94])) = uint8(v561)
	goto L149
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[95])) = v18 + int32(32)
	F_sigprocmask(m, int32(_a_F_AutoVacLauncherMain_9), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L6
	} else {
		goto L169
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[96])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	F_LWLockReleaseAll(m)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[97]))
	*(*int32)(unsafe.Add(mBase, uint32(v717))) = int32(0)
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	F_UnlockBuffers(m)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[98]))
	if v725 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[48]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[47])) = v739
	F_FlushErrorState(m)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[48]))
	F_MemoryContextReset(m, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v747 = int32(_a_F_AutoVacLauncherMain_8)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[49]))
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[49])) = v749 - int32(1)
	v754 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99])) = v754
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100])) = v754
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[101])) = int32(0)
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[102]))
	if v763 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	F_pg_usleep(m, int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L6
	} else {
		goto L168
	}
L167:
	;
	goto L3
L168:
	;
	goto L148
L169:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_11), int32(_a_F_AutoVacLauncherMain_12), int32(5), int32(10))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_13), int32(_a_F_AutoVacLauncherMain_14), int32(5), int32(10))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_15), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_17), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_18), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_19), int32(_a_F_AutoVacLauncherMain_16), int32(5), int32(10))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_20), int32(_a_F_AutoVacLauncherMain_21), int32(5), int32(10))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	F_SetConfigOption(m, int32(_a_F_AutoVacLauncherMain_22), int32(_a_F_AutoVacLauncherMain_23), int32(5), int32(10))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[103])))
	if v826 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[105]))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+8)) = v845
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L6
	} else {
		goto L188
	}
L179:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[106])))
	if v830&int32(1) != 0 {
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[102]))
	if v834 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	v837 = F_do_start_worker(m)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L6
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	goto L3
L188:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[102]))
	if v851 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	goto L192
L190:
	;
	goto L191
L191:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L6
	} else {
		goto L368
	}
L192:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+20))
	v867 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[107]))
	v869 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[108]))
	v870 = v867 - v869
	v871 = int32(0)
	if v871 < v870 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L191
L194:
	;
	v939 = base.I64_extend_i32_s(v937)
	if v937 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L195:
	;
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[109]))
	v937 = v934
	v938 = int32(0)
	goto L194
L196:
	;
	v874 = v870
	goto L198
L197:
	;
	v874 = v871
	goto L198
L198:
	;
	v875 = base.B2i32(v874 < v865)
	if v875 == int32(0) {
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if base.B2i32(v879 == int32(0))|base.B2i32(v879 == int32(_a_F_AutoVacLauncherMain_10)) != 0 {
		goto L195
	} else {
		goto L200
	}
L200:
	;
	v888 = m.G0
	v889 = int32(16)
	v890 = v888 - v889
	m.G0 = v890
	F_gettimeofday(m, v890)
	mBase = m.M
	v893 = *(*int64)(unsafe.Add(mBase, uint32(v890)))
	v894 = int64(*(*int32)(unsafe.Add(mBase, uint32(v890)+8)))
	m.G0 = v890 + v889
	goto L201
L201:
	;
	v904 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100]))
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v904-int32(12))))
	v914 = v907 - (v894 + v893*int64(1000000) - int64(946684800000000))
	if v914 <= int64(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v18)+196))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	v937 = v930
	v938 = v931
	goto L194
L203:
	;
	v926 = int32(0)
	v927 = int32(0)
	goto L205
L204:
	;
	v918 = int64(1000000)
	v919 = base.I64_div_u_s(v914, v918)
	v926 = base.I32_wrap_i64(v919)
	v927 = base.I32_wrap_i64(v914 - v919*v918)
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(196)))) = v926
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(192)))) = v927
	goto L202
L206:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[110]))
	v1042 = int32(1000)
	v1045 = base.I32_div_s(v1034, v1042)
	v1048 = F_WaitLatch(m, v1039, int32(41), base.I32_wrap_i64(v1037)*v1042+v1045, int32(83886081))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L6
	} else {
		goto L245
	}
L207:
	;
	v1030 = int64(300)
	if base.Ui64(v1030) <= base.Ui64(v939) {
		goto L242
	} else {
		goto L243
	}
L208:
	;
	v1026 = base.B2i32(v938 < int32(_a_F_AutoVacLauncherMain_24))
	if v938 < int32(_a_F_AutoVacLauncherMain_24) {
		goto L236
	} else {
		goto L237
	}
L209:
	;
	if v938 != 0 {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	if int32(0) < v937 {
		goto L207
	} else {
		goto L235
	}
L212:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L6
	} else {
		goto L213
	}
L213:
	;
	if v875 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v1011 = base.I64_extend_i32_s(v1010)
	if v1010 <= int32(0) {
		goto L223
	} else {
		goto L224
	}
L215:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[109]))
	v1008 = int32(0)
	v1009 = int32(1)
	v1010 = v1007
	goto L214
L216:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if base.B2i32(v948 == int32(0))|base.B2i32(v948 == int32(_a_F_AutoVacLauncherMain_10)) != 0 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v957 = m.G0
	v958 = int32(16)
	v959 = v957 - v958
	m.G0 = v959
	F_gettimeofday(m, v959)
	mBase = m.M
	v962 = *(*int64)(unsafe.Add(mBase, uint32(v959)))
	v963 = int64(*(*int32)(unsafe.Add(mBase, uint32(v959)+8)))
	m.G0 = v959 + v958
	goto L218
L218:
	;
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100]))
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v973-int32(12))))
	v983 = v976 - (v963 + v962*int64(1000000) - int64(946684800000000))
	if v983 <= int64(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	v1008 = v999
	v1009 = base.B2i32(v999 < int32(_a_F_AutoVacLauncherMain_24))
	v1010 = v1002
	goto L214
L220:
	;
	v995 = int32(0)
	v996 = int32(0)
	goto L222
L221:
	;
	v987 = int64(1000000)
	v988 = base.I64_div_u_s(v983, v987)
	v995 = base.I32_wrap_i64(v988)
	v996 = base.I32_wrap_i64(v983 - v988*v987)
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(204)))) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(200)))) = v996
	goto L219
L223:
	;
	if v1009 != 0 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v1018 = int64(300)
	if base.Ui64(v1018) <= base.Ui64(v1011) {
		goto L232
	} else {
		goto L233
	}
L226:
	;
	v1015 = int64(0)
	goto L228
L227:
	;
	v1015 = v1011
	goto L228
L228:
	;
	if v1009 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1017 = int32(_a_F_AutoVacLauncherMain_25)
	goto L231
L230:
	;
	v1017 = v1008
	goto L231
L231:
	;
	v1034 = v1017
	v1037 = v1015
	goto L206
L232:
	;
	v1021 = v1018
	goto L234
L233:
	;
	v1021 = v1011
	goto L234
L234:
	;
	v1034 = v1008
	v1037 = v1021
	goto L206
L235:
	;
	goto L208
L236:
	;
	v1027 = int64(0)
	goto L238
L237:
	;
	v1027 = v939
	goto L238
L238:
	;
	if v938 < int32(_a_F_AutoVacLauncherMain_24) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1029 = int32(_a_F_AutoVacLauncherMain_25)
	goto L241
L240:
	;
	v1029 = v938
	goto L241
L241:
	;
	v1034 = v1029
	v1037 = v1027
	goto L206
L242:
	;
	v1033 = v1030
	goto L244
L243:
	;
	v1033 = v939
	goto L244
L244:
	;
	v1034 = v938
	v1037 = v1033
	goto L206
L245:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[110]))
	*(*int32)(unsafe.Add(mBase, uint32(v1051))) = int32(0)
	goto L246
L246:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[102]))
	if v1055 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L6
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[111]))
	if v1059 != 0 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L3
L251:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[111])) = int32(0)
	v1064 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[108]))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[112]))
	if v1122 != 0 {
		goto L271
	} else {
		goto L272
	}
L254:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[103])))
	if v1069 == int32(1) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[108]))
	if v1064 == v1079 {
		goto L261
	} else {
		goto L262
	}
L256:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[106])))
	if v1073&int32(1) != 0 {
		goto L255
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L6
	} else {
		goto L260
	}
L259:
	;
	goto L258
L260:
	;
	goto L3
L261:
	;
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L6
	} else {
		goto L270
	}
L262:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[107]))
	if v1079 <= v1082 {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1086 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L6
	} else {
		goto L264
	}
L264:
	;
	if v1086 == int32(0) {
		goto L261
	} else {
		goto L265
	}
L265:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L6
	} else {
		goto L266
	}
L266:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[108]))
	v1096 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[107]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1094
	F_errmsg(m, int32(_a_F_AutoVacLauncherMain_26), v18+int32(16))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[107]))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1105
	F_errdetail(m, int32(_a_F_AutoVacLauncherMain_27), v18)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L6
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(3474), int32(_a_F_AutoVacLauncherMain_28))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L6
	} else {
		goto L269
	}
L269:
	;
	goto L261
L270:
	;
	goto L253
L271:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L6
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[113]))
	if v1126 != 0 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L273
L275:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L6
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L6
	} else {
		goto L279
	}
L278:
	;
	goto L277
L279:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[114]))
	if v1132 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[102]))
	if v1531 == int32(0) {
		goto L192
	} else {
		goto L367
	}
L281:
	;
	v1228 = m.G0
	v1229 = int32(16)
	v1230 = v1228 - v1229
	m.G0 = v1230
	F_gettimeofday(m, v1230)
	mBase = m.M
	v1233 = *(*int64)(unsafe.Add(mBase, uint32(v1230)))
	v1234 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1230)+8)))
	m.G0 = v1230 + v1229
	v1242 = v1234 + v1233*int64(1000000) - int64(946684800000000)
	goto L303
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[114])) = int32(0)
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+4))
	if v1140 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	v1146 = F_LWLockAcquire(m, v1142+int32(2816), int32(0))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L6
	} else {
		goto L286
	}
L284:
	;
	v1196 = v1139
	goto L285
L285:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1196)))
	if v1205 == int32(0) {
		goto L281
	} else {
		goto L300
	}
L286:
	;
	v1148 = int32(0)
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+4)) = v1148
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+uint32(_c_F_AutoVacLauncherMain[116])))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+28))
	if v1154 == v1148 {
		v1179 = v1148
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v1179 != v1153 {
		goto L296
	} else {
		goto L297
	}
L288:
	;
	v1158 = v1150 + int32(24)
	if v1154 == v1158 {
		v1179 = v1148
		goto L287
	} else {
		goto L289
	}
L289:
	;
	v1160 = v1154
	v1162 = v1148
	goto L290
L290:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+16))
	if v1169 != 0 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1179 = v1174
	goto L287
L292:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+32))
	v1174 = v1162 + base.B2i32(v1170 != int32(0))
	goto L294
L293:
	;
	v1174 = v1162
	goto L294
L294:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1160)+4))
	if v1175 != v1158 {
		v1160 = v1175
		v1162 = v1174
		goto L290
	} else {
		goto L295
	}
L295:
	;
	goto L291
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+uint32(_c_F_AutoVacLauncherMain[116]))) = v1179
	goto L298
L297:
	;
	goto L298
L298:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	F_LWLockRelease(m, v1189+int32(2816))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L6
	} else {
		goto L299
	}
L299:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v1196 = v1195
	goto L285
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1196))) = int32(0)
	F_pg_usleep(m, int32(_a_F_AutoVacLauncherMain_3))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L6
	} else {
		goto L301
	}
L301:
	;
	F_SendPostmasterSignal(m, int32(5))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L6
	} else {
		goto L302
	}
L302:
	;
	goto L280
L303:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	v1248 = F_LWLockAcquire(m, v1244+int32(2816), int32(1))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L6
	} else {
		goto L304
	}
L304:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[108]))
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[107]))
	v1255 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+20))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+32))
	if v1257 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	F_LWLockRelease(m, v1341+int32(2816))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L6
	} else {
		goto L322
	}
L306:
	;
	v1260 = *(*int64)(unsafe.Add(mBase, uint32(v1257)+24))
	v1261 = int32(60)
	v1263 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[109]))
	if v1261 <= v1263 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1266 = v1261
	goto L309
L308:
	;
	v1266 = v1263
	goto L309
L309:
	;
	goto L310
L310:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	F_LWLockRelease(m, v1275+int32(2816))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L6
	} else {
		goto L311
	}
L311:
	;
	if base.B2i32(base.I64_extend_i32_s(v1266*int32(1000))*int64(1000) <= v1242-v1260) == int32(0) {
		goto L280
	} else {
		goto L312
	}
L312:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[115]))
	v1287 = F_LWLockAcquire(m, v1283+int32(2816), int32(0))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L6
	} else {
		goto L313
	}
L313:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[104]))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+32))
	if v1291 == int32(0) {
		goto L305
	} else {
		goto L314
	}
L314:
	;
	v1294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1291)+36)) = uint8(v1294)
	v1296 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1291)+8)) = v1296
	*(*int64)(unsafe.Add(mBase, uint32(v1291)+24)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+16)) = v1294
	v1303 = v1290 + int32(12)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+16))
	if v1304 == v1294 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+12)) = v1303
	v1310 = v1303
	goto L317
L316:
	;
	v1310 = v1304
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1291))) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v1291)+4)) = v1310
	*(*int32)(unsafe.Add(mBase, uint32(v1310))) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+16)) = v1291
	v1315 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+32)) = v1315
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1290)+20)) = v1317 + int32(1)
	v1323 = F_errstart(m, int32(19), v1315)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L6
	} else {
		goto L318
	}
L318:
	;
	if v1323 == int32(0) {
		goto L305
	} else {
		goto L319
	}
L319:
	;
	F_errmsg(m, int32(_a_F_AutoVacLauncherMain_29), int32(0))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L6
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_AutoVacLauncherMain_1), int32(693), int32(_a_F_AutoVacLauncherMain_2))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L6
	} else {
		goto L321
	}
L321:
	;
	goto L305
L322:
	;
	v1346 = v1253 - v1251
	v1347 = int32(0)
	if v1347 < v1346 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1350 = v1346
	goto L325
L324:
	;
	v1350 = v1347
	goto L325
L325:
	;
	if v1256 <= v1350 {
		goto L280
	} else {
		goto L326
	}
L326:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if v1353 != int32(_a_F_AutoVacLauncherMain_10) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1357 = v1353
	goto L329
L328:
	;
	v1357 = int32(0)
	goto L329
L329:
	;
	if v1357 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1360 = F_do_start_worker(m)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L6
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100]))
	v1438 = *(*int64)(unsafe.Add(mBase, uint32(v1435-int32(12))))
	goto L349
L333:
	;
	if v1360 == int32(0) {
		goto L280
	} else {
		goto L334
	}
L334:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	v1366 = int32(0)
	if base.B2i32(v1365 == v1366)|base.B2i32(v1365 == int32(_a_F_AutoVacLauncherMain_10)) == v1366 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1373 = v1365
	goto L338
L336:
	;
	goto L337
L337:
	;
	F_rebuild_database_list(m, v1360)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L6
	} else {
		goto L348
	}
L338:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1373-int32(20))))
	if v1360 == v1384 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	goto L337
L340:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[109]))
	*(*int64)(unsafe.Add(mBase, uint32(v1373-int32(12)))) = base.I64_extend_i32_s(v1389*int32(1000))*int64(1000) + v1242
	v1398 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if v1398 == v1373 {
		goto L280
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	if v1420 != int32(_a_F_AutoVacLauncherMain_10) {
		v1373 = v1420
		goto L338
	} else {
		goto L347
	}
L343:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1400)+4)) = v1401
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	*(*int32)(unsafe.Add(mBase, uint32(v1401))) = v1403
	v1406 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if v1406 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1409 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100])) = v1409
	v1413 = v1409
	goto L346
L345:
	;
	v1413 = v1406
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1373))) = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+4)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v1413))) = v1373
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99])) = v1373
	goto L280
L347:
	;
	goto L339
L348:
	;
	goto L280
L349:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(0))*int64(1000) <= v1242-v1438) == int32(0) {
		goto L280
	} else {
		goto L350
	}
L350:
	;
	v1447 = F_do_start_worker(m)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L6
	} else {
		goto L351
	}
L351:
	;
	if v1447 == int32(0) {
		goto L280
	} else {
		goto L352
	}
L352:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	v1453 = int32(0)
	if base.B2i32(v1452 == v1453)|base.B2i32(v1452 == int32(_a_F_AutoVacLauncherMain_10)) == v1453 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1460 = v1452
	goto L356
L354:
	;
	goto L355
L355:
	;
	F_rebuild_database_list(m, v1447)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L6
	} else {
		goto L366
	}
L356:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1460-int32(20))))
	if v1447 == v1471 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L355
L358:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[109]))
	*(*int64)(unsafe.Add(mBase, uint32(v1460-int32(12)))) = base.I64_extend_i32_s(v1476*int32(1000))*int64(1000) + v1242
	v1485 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if v1485 == v1460 {
		goto L280
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+4))
	if v1507 != int32(_a_F_AutoVacLauncherMain_10) {
		v1460 = v1507
		goto L356
	} else {
		goto L365
	}
L361:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1460)))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1487)+4)) = v1488
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1460)))
	*(*int32)(unsafe.Add(mBase, uint32(v1488))) = v1490
	v1493 = *(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99]))
	if v1493 == int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1496 = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[100])) = v1496
	v1500 = v1496
	goto L364
L363:
	;
	v1500 = v1493
	goto L364
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1460))) = int32(_a_F_AutoVacLauncherMain_10)
	*(*int32)(unsafe.Add(mBase, uint32(v1460)+4)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(v1500))) = v1460
	*(*int32)(unsafe.Add(mBase, _c_F_AutoVacLauncherMain[99])) = v1460
	goto L280
L365:
	;
	goto L357
L366:
	;
	goto L280
L367:
	;
	goto L193
L368:
	;
	goto L5
L369:
	;
	v1559 = int32(v1555)
	m.G0 = v18
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1559)+4))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1559)))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1562)))
	if v18+int32(28) == v1565 {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	m.ExcPending = 1
	goto L378
L371:
	;
	if v1569 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L372:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+4))
	v1569 = v1567
	goto L374
L373:
	;
	v1569 = int32(0)
	goto L374
L374:
	;
	goto L371
L375:
	;
	F___wasm_longjmp(m, v1562, v1561)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L378
	} else {
		goto L379
	}
L376:
	;
	goto L377
L377:
	;
	v15 = v1569
	v17 = v1561
	goto L1
L378:
	;
	return
L379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AuxiliaryProcKill(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
	if v6 == int32(42) {
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
					*(*int32)(unsafe.Add(mBase, uint32(v22+int32(20))+12)) = v24
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(1)
					if v32 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
						F_s_lock(m, v36, int32(_a_F_AuxiliaryProcKill_2), int32(1071), int32(_a_F_AuxiliaryProcKill_3))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v22)+52)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[5]))
							v55 = base.I32_div_s(v50+v48*int32(15), int32(16))
							v57 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v57)+68)) = v55
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
							return
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v22)+52)) = int64(4294967295)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[5]))
						v55 = base.I32_div_s(v50+v48*int32(15), int32(16))
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[4]))
						*(*int32)(unsafe.Add(mBase, uint32(v57)+68)) = v55
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcKill[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_AuxiliaryProcKill_4), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_AuxiliaryProcKill_2), int32(1050), int32(_a_F_AuxiliaryProcKill_3))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
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
func F_aclitem_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 != v7 {
		v15 = v2
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v9 != v10 {
			v15 = v2
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v15 = base.B2i32(v12 == v13)
		}
	}
	return v15
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
	F_errfinish(m, int32(_a_F_aclmask_1), int32(1403), int32(_a_F_aclmask_2))
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
func F_aclremove(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_aclremove_0), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_aclremove_1), int32(1607), int32(_a_F_aclremove_2))
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
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
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v103 = F_pg_reg_getnumoutarcs(m, v101, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v107 = F_palloc(m, v103<<(uint(int32(3))%32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_pg_reg_getoutarcs(m, v109, v110, v107, v103)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	if int32(0) < v103 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v122 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v107)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L26
	} else {
		goto L75
	}
L39:
	;
	v126 = int32(-4)
	v130 = v107 + v122<<(uint(int32(3))%32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+40)))
	if v131 != v134 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L38
L41:
	;
	v223 = v122 + int32(1)
	if v223 != v103 {
		v122 = v223
		goto L39
	} else {
		goto L74
	}
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v209 = F_palloc(m, int32(12))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L26
	} else {
		goto L72
	}
L43:
	;
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+42)))
	v139 = base.B2i32(v131 == v136)
	goto L45
L44:
	;
	v139 = int32(1)
	goto L45
L45:
	;
	if v139 != 0 {
		v204 = v126
		v205 = v126
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v140 = int32(-3)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+24))
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+44)))
	if v141 != v144 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v143)+46)))
	v149 = base.B2i32(v141 == v146)
	goto L49
L48:
	;
	v149 = int32(1)
	goto L49
L49:
	;
	if v149 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v204 = int32(-3)
	v205 = v140
	goto L42
L51:
	;
	goto L52
L52:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v151 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v204 = int32(-3)
	v205 = v140
	goto L42
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v159 = v156 + v151*int32(12)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v160 != int32(1) {
		v204 = int32(-3)
		v205 = v140
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v163 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v188 <= int32(0) {
		goto L41
	} else {
		goto L66
	}
L58:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v166 + int32(4) {
	case 0:
		goto L61
	case 1:
		goto L59
	default:
		goto L60
	}
L59:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v177 = F_palloc(m, int32(12))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L26
	} else {
		goto L64
	}
L60:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v172 != int32(-4) {
		goto L57
	} else {
		goto L63
	}
L61:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v169 == int32(-4) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v177)+8)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = int64(-12884901892)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v183 = F_lappend(m, v182, v177)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L26
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v183
	goto L57
L66:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v193 + int32(4) {
	case 0:
		goto L68
	case 1:
		v204 = v192
		v205 = v191
		goto L42
	default:
		goto L67
	}
L67:
	;
	v201 = int32(-4)
	if v191 != v201 {
		goto L41
	} else {
		goto L71
	}
L68:
	;
	if v192 != int32(-4) {
		goto L41
	} else {
		goto L69
	}
L69:
	;
	v198 = int32(-4)
	if v191 == v198 {
		v204 = v192
		v205 = v198
		goto L42
	} else {
		goto L70
	}
L70:
	;
	goto L41
L71:
	;
	v204 = v192
	v205 = v201
	goto L42
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v205
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v215 = F_lappend(m, v214, v209)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L26
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v215
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
	if v33 != int32(319) {
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
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+80)))
	v106 = (v47 - v103) << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v45)+84))
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
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
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
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v116
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v45)+84))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v120+v106)))
	v124 = v122
	goto L33
L36:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v45)+84))
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
	F_errfinish(m, int32(_a_F_add_vars_to_targetlist_1), int32(329), int32(_a_F_add_vars_to_targetlist_2))
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l3 != 0 {
		v12 = F_range_cmp_bounds(m, l0, l1, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)))
			if v16 == int32(1) {
				if v12 < int32(0) {
					v22 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v22
					v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v24
					v30 = F_bounds_adjacent(m, l0, v10+int32(24), v10+int32(16))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v40 = int32(-1)
						} else {
							v40 = int32(1)
						}
						v42 = F_range_cmp_bounds(m, l0, l2, l3)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = int32(0)
							v45 = base.B2i32(v44 <= v42)
							if v45&base.B2i32(v40 < v44)|base.B2i32(base.B2i32(v40 <= v44)|v45 == v44) != 0 {
								v83 = int32(0)
								m.G0 = v10 + int32(32)
								return v83
							} else {
								v58 = F_range_cmp_bounds(m, l0, l1, l2)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
									if v60 == int32(1) {
										if v58 < int32(0) {
											v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v66
											v68 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v68
											v72 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												if v72 == int32(0) {
													v83 = int32(-1)
												} else {
													v83 = int32(1)
												}
												m.G0 = v10 + int32(32)
												return v83
											}
										} else {
											v83 = int32(1)
											m.G0 = v10 + int32(32)
											return v83
										}
									} else {
										if v58 <= int32(0) {
											v81 = int32(-1)
										} else {
											v81 = int32(1)
										}
										v83 = v81
										m.G0 = v10 + int32(32)
										return v83
									}
								}
							}
						}
					}
				} else {
					v40 = int32(1)
					v42 = F_range_cmp_bounds(m, l0, l2, l3)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(0)
						v45 = base.B2i32(v44 <= v42)
						if v45&base.B2i32(v40 < v44)|base.B2i32(base.B2i32(v40 <= v44)|v45 == v44) != 0 {
							v83 = int32(0)
							m.G0 = v10 + int32(32)
							return v83
						} else {
							v58 = F_range_cmp_bounds(m, l0, l1, l2)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
								if v60 == int32(1) {
									if v58 < int32(0) {
										v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v66
										v68 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v68
										v72 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											if v72 == int32(0) {
												v83 = int32(-1)
											} else {
												v83 = int32(1)
											}
											m.G0 = v10 + int32(32)
											return v83
										}
									} else {
										v83 = int32(1)
										m.G0 = v10 + int32(32)
										return v83
									}
								} else {
									if v58 <= int32(0) {
										v81 = int32(-1)
									} else {
										v81 = int32(1)
									}
									v83 = v81
									m.G0 = v10 + int32(32)
									return v83
								}
							}
						}
					}
				}
			} else {
				if v12 <= int32(0) {
					v39 = int32(-1)
				} else {
					v39 = int32(1)
				}
				v40 = v39
				v42 = F_range_cmp_bounds(m, l0, l2, l3)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = int32(0)
					v45 = base.B2i32(v44 <= v42)
					if v45&base.B2i32(v40 < v44)|base.B2i32(base.B2i32(v40 <= v44)|v45 == v44) != 0 {
						v83 = int32(0)
						m.G0 = v10 + int32(32)
						return v83
					} else {
						v58 = F_range_cmp_bounds(m, l0, l1, l2)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
							if v60 == int32(1) {
								if v58 < int32(0) {
									v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v66
									v68 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
									*(*int64)(unsafe.Add(mBase, uint32(v10))) = v68
									v72 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if v72 == int32(0) {
											v83 = int32(-1)
										} else {
											v83 = int32(1)
										}
										m.G0 = v10 + int32(32)
										return v83
									}
								} else {
									v83 = int32(1)
									m.G0 = v10 + int32(32)
									return v83
								}
							} else {
								if v58 <= int32(0) {
									v81 = int32(-1)
								} else {
									v81 = int32(1)
								}
								v83 = v81
								m.G0 = v10 + int32(32)
								return v83
							}
						}
					}
				}
			}
		}
	} else {
		v58 = F_range_cmp_bounds(m, l0, l1, l2)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v60 == int32(1) {
				if v58 < int32(0) {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v66
					v68 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v68
					v72 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						if v72 == int32(0) {
							v83 = int32(-1)
						} else {
							v83 = int32(1)
						}
						m.G0 = v10 + int32(32)
						return v83
					}
				} else {
					v83 = int32(1)
					m.G0 = v10 + int32(32)
					return v83
				}
			} else {
				if v58 <= int32(0) {
					v81 = int32(-1)
				} else {
					v81 = int32(1)
				}
				v83 = v81
				m.G0 = v10 + int32(32)
				return v83
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
func F_alloc_chromo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = F_palloc(m, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v12 = F_palloc(m, l0<<(uint(int32(2))%32)+int32(4))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v12
			return v4
		}
	}
}
func F_anyarray_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_anycompatiblearray_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_anycompatiblearray_recv_0), int32(175), int32(_a_F_anycompatiblearray_recv_1), int32(_a_F_anycompatiblearray_recv_2), int32(_a_F_anycompatiblearray_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13845(m, l0, l1, int32(_a_F_anytime_typmod_check_0), int32(78), int32(_a_F_anytime_typmod_check_1), int32(_a_F_anytime_typmod_check_2), int32(85), int32(_a_F_anytime_typmod_check_3))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(1200)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v18 = F_LWLockAcquire(m, v16, v3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	if v24 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v339 = int32(_a_F_apw_dump_now_0)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v342 = v13 + int32(176)
	v343 = F_unlink(m, v342)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v340
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L78
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L74
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L71
	}
L6:
	;
	m.G0 = v13 + int32(1200)
	return v294
L7:
	;
	F_LWLockRelease(m, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v52
	F_LWLockRelease(m, v23)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v33 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v33 == int32(0) {
		v294 = v3
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v39
	F_errmsg(m, int32(_a_F_apw_dump_now_1), v13+int32(144))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(692), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v294 = v3
	goto L6
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	v61 = F_palloc_extended(m, v57*int32(20), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	if int32(0) < v64 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = int32(0)
	v72 = v3
	goto L21
L19:
	;
	v130 = v3
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = int32(_a_F_apw_dump_now_4)
	v140 = v13 + int32(176)
	v145 = F_pg_snprintf(m, v140, int32(1024), int32(_a_F_apw_dump_now_5), v13+int32(128))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[4]))
	if v79 != 0 {
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
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[5]))
	v86 = v83 + v68<<(uint(int32(6))%32)
	v87 = F_LockBufHdr(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v91 = int32(0)
	if base.B2i32(v87&int32(33554432) == v91)|base.B2i32(l1|base.B2i32(v87 < v91) == v91) == v91 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v103 = v61 + v72*int32(20)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v112
	v116 = v72 + int32(1)
	goto L30
L29:
	;
	v116 = v72
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+24)) = v87 & int32(-4194305)
	v122 = v68 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[3]))
	if v122 < v124 {
		v68 = v122
		v72 = v116
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L22
L32:
	;
	v148 = F_AllocateFile(m, v140, int32(_a_F_apw_dump_now_6))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v148 == int32(0) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v130
	v156 = F_pg_fprintf(m, v148, int32(_a_F_apw_dump_now_7), v13+int32(112))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	F_pfree(m, v61)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L63
	}
L36:
	;
	v191 = int32(0)
	goto L47
L37:
	;
	if int32(0) <= v156 {
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
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v166 = F_FreeFile(m, v148)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L36
L42:
	;
	v169 = v13 + int32(176)
	v170 = F_unlink(m, v169)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v165
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v169
	F_errmsg(m, int32(_a_F_apw_dump_now_8), v13+int32(16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(757), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
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
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[4]))
	if v201 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1]))
	v225 = F_FreeFile(m, v148)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L58
	}
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v206 = v61 + v191*int32(20)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v206)))
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v206)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(96)))) = v209
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v208
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v207
	v216 = F_pg_fprintf(m, v148, int32(_a_F_apw_dump_now_9), v13+int32(80))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	if int32(0) <= v216 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v221 = v191 + int32(1)
	if v221 == v130 {
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
	v191 = v221
	goto L47
L58:
	;
	v228 = v13 + int32(176)
	v229 = F_unlink(m, v228)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[1])) = v224
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v228
	F_errmsg(m, int32(_a_F_apw_dump_now_8), v13+int32(32))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(780), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
	v261 = F_FreeFile(m, v148)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v261 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v267 = F_durable_rename(m, v13+int32(176), int32(_a_F_apw_dump_now_4), int32(21))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+20)) = int32(-1)
	v275 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v275 == int32(0) {
		v294 = v130
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v130
	F_errmsg_internal(m, int32(_a_F_apw_dump_now_10), v13+int32(48))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(807), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v294 = v130
	goto L6
L71:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_apw_dump_now[0]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+160)) = v310
	F_errmsg(m, int32(_a_F_apw_dump_now_11), v13+int32(160))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(688), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
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
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v13 + int32(176)
	F_errmsg(m, int32(_a_F_apw_dump_now_12), v13)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(744), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v342
	F_errmsg(m, int32(_a_F_apw_dump_now_13), v13-int32(-64))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_apw_dump_now_2), int32(800), int32(_a_F_apw_dump_now_3))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v7 = m.G0
	v9 = v7 - int32(1472)
	m.G0 = v9
	base.MemoryFill(m, v9+int32(24), int32(0), int32(1444))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+200)) = int64(4294967297)
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+212)) = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+219)) = v22
	v25 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+1236)) = v25
	v28 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+1244)) = v28
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[4])))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+1252)) = uint8(v31)
	v34 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v37
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+23)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v9)+119)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v34
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[8])))
	if v46 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L35
	}
L3:
	;
	m.G0 = v9 + int32(1472)
	return
L4:
	;
	F_RegisterBackgroundWorker(m, v9+int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+1464)) = v54
	v60 = F_RegisterDynamicBackgroundWorker(m, v9+int32(8), v9+int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
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
	if v60 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	goto L12
L11:
	;
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[10]))
	if v72 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v102
	v129 = int32(0)
	goto L11
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[11]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	v83 = F_LWLockAcquire(m, v79+int32(_a_F_apw_start_leader_worker_0), int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
	v88 = v76 + v77*int32(1480)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)+24))
	if v85 == v89 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	F_LWLockRelease(m, v104+int32(_a_F_apw_start_leader_worker_0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L25
	}
L20:
	;
	v92 = v88 + int32(16)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != 0 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[12]))
	F_LWLockRelease(m, v96+int32(_a_F_apw_start_leader_worker_0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v129 = int32(2)
	goto L11
L25:
	;
	if v102 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L13
L27:
	;
	if v102 != 0 {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[13]))
	v118 = F_WaitLatch(m, v114, int32(17), int32(0), int32(134217734))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	v129 = int32(2)
	goto L11
L31:
	;
	if v118&int32(16) != 0 {
		v129 = int32(3)
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_apw_start_leader_worker[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = int32(0)
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_apw_start_leader_worker_1), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_errhint(m, int32(_a_F_apw_start_leader_worker_2), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_apw_start_leader_worker_3), int32(936), int32(_a_F_apw_start_leader_worker_4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_apw_start_leader_worker_5), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_errhint(m, int32(_a_F_apw_start_leader_worker_6), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_apw_start_leader_worker_3), int32(943), int32(_a_F_apw_start_leader_worker_4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
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
	var v11 int32
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
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
							F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1154), int32(_a_F_assignOperTypes_2))
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
											F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1174), int32(_a_F_assignOperTypes_2))
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
									F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1184), int32(_a_F_assignOperTypes_2))
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
					F_errfinish(m, int32(_a_F_assignOperTypes_1), int32(1145), int32(_a_F_assignOperTypes_2))
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
func F_assign_io_method(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_assign_io_method[0])))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_io_method[1])) = v8
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
