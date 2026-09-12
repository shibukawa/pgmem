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
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 float64
	_ = v77
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v100 int32
	_ = v100
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
		v100 = v5
	} else {
		v58 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
		v59 = v58 + v54
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v59
		if base.B2i32(v54 < int64(0))^base.B2i32(v59 < v58) != 0 {
			v100 = v5
		} else {
			if base.F64_eq(l1, float64(0)) != 0 {
				v100 = int32(1)
			} else {
				v69 = base.F64_mul(l1, base.F64_convert_i64_u(l2))
				if base.F64_lt(base.F64_abs(v69), float64(9.223372036854776e+18)) != 0 {
					v73 = base.I64_trunc_f64_s(v69)
					v75 = v73
				} else {
					v75 = int64(-9223372036854775807 - 1)
				}
				v77 = base.F64_sub(v69, base.F64_convert_i64_s(v75))
				if base.F64_gt(v77, float64(0.5)) != 0 {
					v88 = v75 + int64(1)
				} else {
					if base.F64_lt(v77, float64(-0.5)) == int32(0) {
						v88 = v75
					} else {
						v88 = v75 - int64(1)
					}
				}
				v89 = v88 + v59
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v89
				v100 = base.B2i32(base.B2i32(v88 < int64(0))^base.B2i32(v89 < v59) == int32(0))
			}
		}
	}
	m.G0 = v11 + int32(16)
	return v100
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		v18 = base.B2i32(base.Ui32(v6) <= base.Ui32(l0))
	} else {
		v18 = base.B2i32(int32(0) <= l0-v6)
	}
	if v18 != 0 {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[171]))
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
		v35 = *(*int32)(unsafe.Add(mBase, _consts[86]))
		v39 = F_LWLockAcquire(m, v35+int32(384), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, _consts[171]))
			*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = base.I64_extend_i32_u(v31) | v33<<(uint(int64(32))%64)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
		v29 = v5 << (uint(int32(3)) % 32)
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v32 == int32(0) {
			F_MemoryContextResetOnly(m, l0)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
				if v37 < int32(100) {
					v57 = v37
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441])))
					if v40 == int32(0) {
						v57 = v37
					} else {
						v46 = v40
						for {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441]))) = v47
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
							*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440]))) = v49 - int32(1)
							F_emscripten_builtin_free(m, v46)
							mBase = m.M
							if v47 != 0 {
								v46 = v47
								continue
							} else {
								break
							}
							break
						}
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
						v57 = v54
					}
				}
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441])))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440]))) = v57 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441]))) = l0
				return
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
			if v37 < int32(100) {
				v57 = v37
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441])))
				if v40 == int32(0) {
					v57 = v37
				} else {
					v46 = v40
					for {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441]))) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440]))) = v49 - int32(1)
						F_emscripten_builtin_free(m, v46)
						mBase = m.M
						if v47 != 0 {
							v46 = v47
							continue
						} else {
							break
						}
						break
					}
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440])))
					v57 = v54
				}
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441])))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59
			*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1440]))) = v57 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1441]))) = l0
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
											v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
									v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
				v21 = *(*int32)(unsafe.Add(mBase, _consts[159]))
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
							v33 = *(*int32)(unsafe.Add(mBase, _consts[159]))
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
																	v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
															v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
																			v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
																	v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
						v33 = *(*int32)(unsafe.Add(mBase, _consts[159]))
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
																v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
														v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
																		v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
																v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
		v93 = *(*int32)(unsafe.Add(mBase, _consts[380]))
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
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
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
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
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
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
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
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int64
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v385 = int32(2)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v380-int32(1023)) < base.Ui32(v385) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v380 = v11
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v370
	v380 = v370
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
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v370 = v365 + int32(4)
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
	v370 = v361
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
	v355 = int32(1024) - v352
	if base.Ui32(v25) < base.Ui32(v355) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v347 = F_Int64GetDatum(m, base.I64_extend_i32_u(v337)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v329^v337-base.I32_rotl(v337, int32(24))))
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
	v357 = v25
	goto L61
L60:
	;
	v357 = v355
	goto L61
L61:
	;
	if v357 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v361 = v352 + v357
	v362 = v25 - v357
	if v362 != 0 {
		v23 = v361
		v25 = v362
		v27 = v357 + v27
		goto L9
	} else {
		goto L66
	}
L63:
	;
	v358 = F__emscripten_memcpy_bulkmem(m, v352+v13, v27, v357)
	mBase = m.M
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L10
L67:
	;
	v392 = l1
	v393 = v380
	v396 = v385
	goto L70
L68:
	;
	goto L69
L69:
	;
	v735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*uint16)(unsafe.Add(mBase, uint32(v380+v386))) = uint16(v735)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v737 + int32(2)
	return
L70:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v393) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v731
	return
L72:
	;
	v400 = int32(1024)
	v407 = int32(-1636607408)
	goto L77
L73:
	;
	v722 = v393
	goto L74
L74:
	;
	v725 = int32(1024) - v722
	if base.Ui32(v396) < base.Ui32(v725) {
		goto L119
	} else {
		goto L120
	}
L75:
	;
	v717 = F_Int64GetDatum(m, base.I64_extend_i32_u(v707)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v699^v707-base.I32_rotl(v707, int32(24))))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L57
	} else {
		goto L118
	}
L76:
	;
	if v386&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L77:
	;
	goto L76
L80:
	;
	v685 = int32(14)
	v687 = v681 ^ v682 - base.I32_rotl(v681, v685)
	v691 = v687 ^ v680 - base.I32_rotl(v687, int32(11))
	v695 = v691 ^ v681 - base.I32_rotl(v691, int32(25))
	v699 = v695 ^ v687 - base.I32_rotl(v695, int32(16))
	v703 = v699 ^ v691 - base.I32_rotl(v699, int32(4))
	v707 = v703 ^ v695 - base.I32_rotl(v703, v685)
	goto L75
L81:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v680 = v672 + v675
	v681 = v673
	v682 = v674
	goto L80
L82:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)))
	v672 = v668<<(uint(int32(8))%32) + v665
	v673 = v666
	v674 = v667
	goto L81
L83:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+2)))
	v665 = v661<<(uint(int32(16))%32) + v658
	v666 = v659
	v667 = v660
	goto L82
L84:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+3)))
	v658 = v654<<(uint(int32(24))%32) + v490
	v659 = v652
	v660 = v653
	goto L83
L85:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+4)))
	v652 = v648 + v650
	v653 = v649
	goto L84
L86:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+5)))
	v648 = v644<<(uint(int32(8))%32) + v642
	v649 = v643
	goto L85
L87:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+6)))
	v642 = v638<<(uint(int32(16))%32) + v636
	v643 = v637
	goto L86
L88:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+7)))
	v636 = v632<<(uint(int32(24))%32) + v491
	v637 = v631
	goto L87
L89:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+8)))
	v631 = v627<<(uint(int32(8))%32) + v626
	goto L88
L90:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+9)))
	v626 = v622<<(uint(int32(16))%32) + v621
	goto L89
L91:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+10)))
	v621 = v617<<(uint(int32(24))%32) + v495
	goto L90
L92:
	;
	goto L95
L93:
	;
	goto L94
L94:
	;
	goto L101
L95:
	;
	v453 = v386
	v454 = v400
	v456 = v407
	v457 = v407
	v458 = v407
	goto L98
L97:
	;
	switch v499 - int32(1) {
	case 0:
		v672 = v490
		v673 = v491
		v674 = v495
		goto L81
	case 1:
		v665 = v490
		v666 = v491
		v667 = v495
		goto L82
	case 2:
		v658 = v490
		v659 = v491
		v660 = v495
		goto L83
	case 3:
		v652 = v491
		v653 = v495
		goto L84
	case 4:
		v648 = v491
		v649 = v495
		goto L85
	case 5:
		v642 = v491
		v643 = v495
		goto L86
	case 6:
		v636 = v491
		v637 = v495
		goto L87
	case 7:
		v631 = v495
		goto L88
	case 8:
		v626 = v495
		goto L89
	case 9:
		v621 = v495
		goto L90
	case 10:
		goto L91
	default:
		v680 = v490
		v681 = v491
		v682 = v495
		goto L80
	}
L98:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v461 = v460 + v457
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v465 = v464 + v458
	v467 = int32(4)
	v469 = v462 + v456 - v465 ^ base.I32_rotl(v465, v467)
	v473 = v461 - v469 ^ base.I32_rotl(v469, int32(6))
	v474 = v465 + v461
	v475 = v469 + v474
	v476 = v473 + v475
	v480 = v474 - v473 ^ base.I32_rotl(v473, int32(8))
	v484 = v475 - v480 ^ base.I32_rotl(v480, int32(16))
	v488 = v476 - v484 ^ base.I32_rotl(v484, int32(19))
	v489 = v480 + v476
	v490 = v484 + v489
	v491 = v488 + v490
	v495 = v489 - v488 ^ base.I32_rotl(v488, v467)
	v496 = int32(12)
	v497 = v453 + v496
	v499 = v454 - v496
	if base.Ui32(int32(11)) < base.Ui32(v499) {
		v453 = v497
		v454 = v499
		v456 = v490
		v457 = v491
		v458 = v495
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	goto L99
L101:
	;
	v513 = v386
	v514 = v400
	v516 = v407
	v517 = v407
	v518 = v407
	goto L104
L103:
	;
	switch v559 - int32(1) {
	case 0:
		v614 = v550
		goto L107
	case 1:
		v609 = v550
		goto L108
	case 2:
		goto L109
	case 3:
		v602 = v551
		goto L110
	case 4:
		v599 = v551
		goto L111
	case 5:
		v594 = v551
		goto L112
	case 6:
		goto L113
	case 7:
		v585 = v555
		goto L114
	case 8:
		v580 = v555
		goto L115
	case 9:
		v575 = v555
		goto L116
	case 10:
		goto L117
	default:
		v680 = v550
		v681 = v551
		v682 = v555
		goto L80
	}
L104:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v521 = v520 + v517
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	v525 = v524 + v518
	v527 = int32(4)
	v529 = v522 + v516 - v525 ^ base.I32_rotl(v525, v527)
	v533 = v521 - v529 ^ base.I32_rotl(v529, int32(6))
	v534 = v525 + v521
	v535 = v529 + v534
	v536 = v533 + v535
	v540 = v534 - v533 ^ base.I32_rotl(v533, int32(8))
	v544 = v535 - v540 ^ base.I32_rotl(v540, int32(16))
	v548 = v536 - v544 ^ base.I32_rotl(v544, int32(19))
	v549 = v540 + v536
	v550 = v544 + v549
	v551 = v548 + v550
	v555 = v549 - v548 ^ base.I32_rotl(v548, v527)
	v556 = int32(12)
	v557 = v513 + v556
	v559 = v514 - v556
	if base.Ui32(int32(11)) < base.Ui32(v559) {
		v513 = v557
		v514 = v559
		v516 = v550
		v517 = v551
		v518 = v555
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
L107:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v680 = v614 + v615
	v681 = v551
	v682 = v555
	goto L80
L108:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	v614 = v610<<(uint(int32(8))%32) + v609
	goto L107
L109:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+2)))
	v609 = v605<<(uint(int32(16))%32) + v550
	goto L108
L110:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v680 = v603 + v550
	v681 = v602
	v682 = v555
	goto L80
L111:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+4)))
	v602 = v599 + v600
	goto L110
L112:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+5)))
	v599 = v595<<(uint(int32(8))%32) + v594
	goto L111
L113:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+6)))
	v594 = v590<<(uint(int32(16))%32) + v551
	goto L112
L114:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	v680 = v586 + v550
	v681 = v588 + v551
	v682 = v585
	goto L80
L115:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+8)))
	v585 = v581<<(uint(int32(8))%32) + v580
	goto L114
L116:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+9)))
	v580 = v576<<(uint(int32(16))%32) + v575
	goto L115
L117:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+10)))
	v575 = v571<<(uint(int32(24))%32) + v555
	goto L116
L118:
	;
	v719 = *(*int64)(unsafe.Add(mBase, uint32(v717)))
	*(*int64)(unsafe.Add(mBase, uint32(v386))) = v719
	v722 = int32(8)
	goto L74
L119:
	;
	v727 = v396
	goto L121
L120:
	;
	v727 = v725
	goto L121
L121:
	;
	if v727 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v731 = v722 + v727
	v732 = v396 - v727
	if v732 != 0 {
		v392 = v392 + v727
		v393 = v731
		v396 = v732
		goto L70
	} else {
		goto L126
	}
L123:
	;
	v728 = F__emscripten_memcpy_bulkmem(m, v722+v386, v392, v727)
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	goto L71
}
func F_ApplySetting(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
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
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	F_ScanKeyInit(m, v9+int32(16), int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v9-int32(-64), int32(2), int32(3), int32(184), l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = F_systable_beginscan(m, l3, int32(2965), int32(1), l0, int32(2), v9+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = F_systable_getnext(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = v32
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_systable_endscan(m, v30)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v43 = F_heap_getattr_4(m, v35, v40, v9+int32(15))
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
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
	v54 = F_systable_getnext(m, v30)
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
		v35 = v54
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	m.G0 = v9 + int32(112)
	return
}
func F_AutoVacLauncherMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v244 int64
	_ = v244
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int64
	_ = v341
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v528 int32
	_ = v528
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v917 int32
	_ = v917
	var v925 int32
	_ = v925
	var v933 int32
	_ = v933
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1196 int32
	_ = v1196
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int64
	_ = v1266
	var v1267 int64
	_ = v1267
	var v1275 int64
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int64
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1337 int64
	_ = v1337
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1377 int32
	_ = v1377
	var v1384 int32
	_ = v1384
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1429 int32
	_ = v1429
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1475 int32
	_ = v1475
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1499 int64
	_ = v1499
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1527 int32
	_ = v1527
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1592 int32
	_ = v1592
	var v1607 int32
	_ = v1607
	var v1626 int32
	_ = v1626
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1641 int64
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v21 = v16
	v22 = v3
	v23 = v3
	v24 = int32(-1)
	v25 = v3
	v27 = v16
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
	if v24 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1640 = int32(m.ExcTag)
	v1641 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1640 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L7:
	;
	v35 = v27 - int32(160)
	m.G0 = v35
	v38 = v35 - int32(16)
	m.G0 = v38
	v41 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v628 = v22
	v629 = v23
	v631 = v25
	v632 = v27
	goto L9
L9:
	;
	if v631 != 0 {
		goto L146
	} else {
		goto L147
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_MemoryContextDelete(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[571])) = int32(0)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v64 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L18
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	v59 = F_GetBackendTypeDesc(m, v58)
	mBase = m.M
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_errmsg_internal(m, int32(437509), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[635]))
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_errfinish(m, int32(488518), int32(385), int32(274356))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_pg_usleep(m, v80*int32(1000000))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v90 = int32(913)
	v92 = m.G0
	v94 = v92 - int32(144)
	m.G0 = v94
	switch int32(915) {
	case 0, 2:
		v104 = v90
		goto L29
	default:
		goto L30
	}
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v134 = int32(914)
	v136 = m.G0
	v138 = v136 - int32(144)
	m.G0 = v138
	switch int32(916) {
	case 0, 2:
		v148 = v134
		goto L42
	default:
		goto L43
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v104
	F_sigemptyset(m, v94+int32(8))
	mBase = m.M
	goto L32
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v90
	v104 = int32(4729)
	goto L29
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+136)) = int32(268435456)
	v116 = v94 + int32(4)
	goto L36
L34:
	;
	m.G0 = v94 + int32(144)
	goto L28
L36:
	;
	goto L37
L37:
	;
	if v116 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v127 = F___memcpy(m, int32(4635612), v116, int32(140))
	mBase = m.M
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L34
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v178 = int32(915)
	v180 = m.G0
	v182 = v180 - int32(144)
	m.G0 = v182
	switch int32(917) {
	case 0, 2:
		v192 = v178
		goto L55
	default:
		goto L56
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v148
	F_sigemptyset(m, v138+int32(8))
	mBase = m.M
	goto L45
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v134
	v148 = int32(4729)
	goto L42
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+136)) = int32(268435456)
	v160 = v138 + int32(4)
	goto L49
L47:
	;
	m.G0 = v138 + int32(144)
	goto L41
L49:
	;
	goto L50
L50:
	;
	if v160 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v171 = F___memcpy(m, int32(4635752), v160, int32(140))
	mBase = m.M
	goto L53
L52:
	;
	goto L53
L53:
	;
	goto L47
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[636])) = v221
	*(*int32)(unsafe.Add(mBase, _consts[637])) = v221
	v231 = v221
	goto L68
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v192
	F_sigemptyset(m, v182+int32(8))
	mBase = m.M
	goto L58
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[323])) = v178
	v192 = int32(4729)
	goto L55
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+136)) = int32(268435456)
	v204 = v182 + int32(4)
	goto L62
L60:
	;
	m.G0 = v182 + int32(144)
	goto L54
L62:
	;
	goto L63
L63:
	;
	if v204 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v215 = F___memcpy(m, int32(4637572), v204, int32(140))
	mBase = m.M
	goto L66
L65:
	;
	goto L66
L66:
	;
	goto L60
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v370 = int32(-2)
	v372 = m.G0
	v374 = v372 - int32(144)
	m.G0 = v374
	switch int32(0) {
	case 0, 2:
		v384 = v370
		goto L74
	default:
		goto L75
	}
L68:
	;
	v233 = int32(40)
	v234 = v231 * v233
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[638]))) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[639]))) = v231
	v244 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[640]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[641]))) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[642]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[643]))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[644]))) = uint8(v237)
	v263 = v231 | int32(1)
	v265 = v263 * v233
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[638]))) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[639]))) = v263
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[640]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[641]))) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[642]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[643]))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+uint32(_consts[644]))) = uint8(v237)
	v294 = v231 | int32(2)
	v296 = v294 * v233
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[638]))) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[639]))) = v294
	*(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[640]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[641]))) = v237
	*(*int64)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[642]))) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[643]))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v296)+uint32(_consts[644]))) = uint8(v237)
	if base.B2i32(v231 == int32(20)) == v237 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[645])) = uint8(v362)
	F_pqsignal_be(m, int32(14), int32(1784))
	mBase = m.M
	goto L67
L70:
	;
	v329 = v231 | int32(3)
	v331 = v329 * int32(40)
	v334 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[638]))) = uint8(v334)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[639]))) = v329
	v341 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[640]))) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[641]))) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[642]))) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[643]))) = v334
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+uint32(_consts[644]))) = uint8(v334)
	v231 = v231 + int32(4)
	goto L68
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v414 = int32(916)
	v416 = m.G0
	v418 = v416 - int32(144)
	m.G0 = v418
	switch int32(918) {
	case 0, 2:
		v428 = v414
		goto L87
	default:
		goto L88
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = v384
	F_sigemptyset(m, v374+int32(8))
	mBase = m.M
	goto L77
L75:
	;
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v370
	v384 = int32(4729)
	goto L74
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+136)) = int32(268435456)
	v396 = v374 + int32(4)
	goto L81
L79:
	;
	m.G0 = v374 + int32(144)
	goto L73
L81:
	;
	goto L82
L82:
	;
	if v396 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v407 = F___memcpy(m, int32(4637292), v396, int32(140))
	mBase = m.M
	goto L85
L84:
	;
	goto L85
L85:
	;
	goto L79
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v458 = int32(917)
	v460 = m.G0
	v462 = v460 - int32(144)
	m.G0 = v462
	switch int32(919) {
	case 0, 2:
		v472 = v458
		goto L100
	default:
		goto L101
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v428
	F_sigemptyset(m, v418+int32(8))
	mBase = m.M
	goto L90
L88:
	;
	*(*int32)(unsafe.Add(mBase, _consts[647])) = v414
	v428 = int32(4729)
	goto L87
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+136)) = int32(268435456)
	v440 = v418 + int32(4)
	goto L94
L92:
	;
	m.G0 = v418 + int32(144)
	goto L86
L94:
	;
	goto L95
L95:
	;
	if v440 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v451 = F___memcpy(m, int32(4636872), v440, int32(140))
	mBase = m.M
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L92
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v502 = int32(918)
	v504 = m.G0
	v506 = v504 - int32(144)
	m.G0 = v506
	switch int32(920) {
	case 0, 2:
		v516 = v502
		goto L113
	default:
		goto L114
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v472
	F_sigemptyset(m, v462+int32(8))
	mBase = m.M
	goto L103
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v458
	v472 = int32(4729)
	goto L100
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+136)) = int32(268435456)
	v484 = v462 + int32(4)
	goto L107
L105:
	;
	m.G0 = v462 + int32(144)
	goto L99
L107:
	;
	goto L108
L108:
	;
	if v484 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v495 = F___memcpy(m, int32(4637152), v484, int32(140))
	mBase = m.M
	goto L111
L110:
	;
	goto L111
L111:
	;
	goto L105
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v546 = int32(0)
	v548 = m.G0
	v550 = v548 - int32(144)
	m.G0 = v550
	switch int32(2) {
	case 0, 2:
		v560 = v546
		goto L126
	default:
		goto L127
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v516
	F_sigemptyset(m, v506+int32(8))
	mBase = m.M
	goto L116
L114:
	;
	*(*int32)(unsafe.Add(mBase, _consts[649])) = v502
	v516 = int32(4729)
	goto L113
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+136)) = int32(268435456)
	v528 = v506 + int32(4)
	goto L120
L118:
	;
	m.G0 = v506 + int32(144)
	goto L112
L120:
	;
	goto L121
L121:
	;
	if v528 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v539 = F___memcpy(m, int32(4636592), v528, int32(140))
	mBase = m.M
	goto L124
L123:
	;
	goto L124
L124:
	;
	goto L118
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_InitProcess(m)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L138
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+4)) = v560
	F_sigemptyset(m, v550+int32(8))
	mBase = m.M
	goto L128
L127:
	;
	*(*int32)(unsafe.Add(mBase, _consts[650])) = v546
	v560 = int32(4729)
	goto L126
L128:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+136)) = int32(268435457)
	v572 = v550 + int32(4)
	goto L133
L131:
	;
	m.G0 = v550 + int32(144)
	goto L125
L133:
	;
	goto L134
L134:
	;
	if v572 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v583 = F___memcpy(m, int32(4637852), v572, int32(140))
	mBase = m.M
	goto L137
L136:
	;
	goto L137
L137:
	;
	goto L131
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	F_BaseInit(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	v597 = int32(0)
	F_InitPostgres(m, v597, v597, v597, v597, v597, v597)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _consts[123])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v35
	v611 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v616 = F_AllocSetContextCreateInternal(m, v611, int32(219991), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		v1635 = v38
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v616
	*(*int32)(unsafe.Add(mBase, _consts[651])) = v616
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v21 + int32(36)
	goto L145
L143:
	;
	v628 = v38
	v629 = v35
	v631 = int32(0)
	v632 = v38
	goto L9
L145:
	;
	goto L143
L146:
	;
	v633 = int32(4465212)
	v635 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v635 + int32(1)
	v640 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[385])) = v640
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, _consts[636])) = v640
	*(*int32)(unsafe.Add(mBase, _consts[637])) = v640
	*(*uint8)(unsafe.Add(mBase, _consts[638])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[644])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[652])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[653])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[654])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[655])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[656])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[657])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[658])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[659])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[660])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[661])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[662])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[663])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[664])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[665])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[666])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[667])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[668])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[669])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[670])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[671])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[672])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[673])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[674])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[675])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[676])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[677])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[678])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[679])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[680])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[681])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[682])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[683])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[684])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[685])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[686])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[687])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[688])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[689])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[690])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[691])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[692])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[693])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[694])) = uint8(v640)
	*(*uint8)(unsafe.Add(mBase, _consts[695])) = uint8(v640)
	goto L149
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	F_sigprocmask(m, int32(4377784), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L169
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	F_EmitErrorReport(m)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_LWLockReleaseAll(m)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v804 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v804))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_UnlockBuffers(m)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	if v816 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_smgrdestroyall(m)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L161
	}
L161:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v840
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	F_FlushErrorState(m)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v849 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	F_MemoryContextReset(m, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v852 = int32(4465212)
	v854 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v854 - int32(1)
	v859 = int32(4088592)
	*(*int32)(unsafe.Add(mBase, _consts[697])) = v859
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v859
	*(*int32)(unsafe.Add(mBase, _consts[699])) = int32(0)
	v868 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	if v868 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_pg_usleep(m, int32(1000000))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		v1635 = v632
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(316254), int32(731620), int32(5), int32(10))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(167865), int32(355269), int32(5), int32(10))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(64659), int32(546158), int32(5), int32(10))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(64750), int32(546158), int32(5), int32(10))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(64854), int32(546158), int32(5), int32(10))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(64793), int32(546158), int32(5), int32(10))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(258042), int32(433682), int32(5), int32(10))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SetConfigOption(m, int32(22698), int32(366336), int32(5), int32(10))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, _consts[374])))
	if v951 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v972 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v974 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	*(*int32)(unsafe.Add(mBase, uint32(v972)+8)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L188
	}
L179:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, _consts[375])))
	if v955&int32(1) != 0 {
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v959 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	if v959 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v964 = F_do_start_worker(m)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_proc_exit(m, int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		v1635 = v632
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
	v982 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	if v982 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L321
	}
L192:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1004 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v1006 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	v1007 = v1004 - v1006
	v1008 = int32(0)
	if v1008 < v1007 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L191
L194:
	;
	v1011 = v1007
	goto L196
L195:
	;
	v1011 = v1008
	goto L196
L196:
	;
	F_launcher_determine_sleep(m, base.B2i32(v1011 < v1000), int32(0), v628)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L197
	}
L197:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v628)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1021 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v1023 = int32(1000)
	v1026 = base.I32_div_s(v1017, v1023)
	v1029 = F_WaitLatch(m, v1021, int32(41), v1016*v1023+v1026, int32(83886081))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1034 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v1034))) = int32(0)
	goto L199
L199:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	if v1038 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v1044 != 0 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L3
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
	v1051 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, _consts[706]))
	if v1123 != 0 {
		goto L224
	} else {
		goto L225
	}
L207:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, _consts[374])))
	if v1056 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	if v1051 == v1068 {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, _consts[375])))
	if v1060&int32(1) != 0 {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_AutoVacLauncherShutdown(m)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	goto L3
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_rebuild_database_list(m, int32(0))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L223
	}
L215:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	if v1068 <= v1071 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v1077 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L217
	}
L217:
	;
	if v1077 == int32(0) {
		goto L214
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1089 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	v1091 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v1091
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1089
	F_errmsg(m, int32(654604), v21+int32(16))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1102 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1102
	F_errdetail(m, int32(609653), v21)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_errfinish(m, int32(488518), int32(3474), int32(171010))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L222
	}
L222:
	;
	goto L214
L223:
	;
	goto L206
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	if v1129 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L226
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_ProcessCatchupInterrupt(m)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L232
	}
L231:
	;
	goto L230
L232:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _consts[708]))
	if v1139 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	if v1607 == int32(0) {
		goto L192
	} else {
		goto L320
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v1261 = m.G0
	v1262 = int32(16)
	v1263 = v1261 - v1262
	m.G0 = v1263
	F___gettimeofday(m, v1263)
	mBase = m.M
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(v1263)))
	v1267 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1263)+8)))
	m.G0 = v1263 + v1262
	v1275 = v1267 + v1266*int64(1000000) - int64(946684800000000)
	goto L256
L235:
	;
	*(*int32)(unsafe.Add(mBase, _consts[708])) = int32(0)
	v1146 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	if v1147 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1151 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1155 = F_LWLockAcquire(m, v1151+int32(2816), int32(0))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L239
	}
L237:
	;
	v1220 = v1146
	goto L238
L238:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	if v1228 == int32(0) {
		goto L234
	} else {
		goto L253
	}
L239:
	;
	v1157 = int32(0)
	v1159 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+4)) = v1157
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+uint32(_consts[709])))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+28))
	if v1163 == v1157 {
		v1196 = v1157
		goto L240
	} else {
		goto L241
	}
L240:
	;
	if v1196 != v1162 {
		goto L249
	} else {
		goto L250
	}
L241:
	;
	v1167 = v1159 + int32(24)
	if v1163 == v1167 {
		v1196 = v1157
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1174 = v1163
	v1175 = v1157
	goto L243
L243:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+16))
	if v1182 != 0 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v1196 = v1187
	goto L240
L245:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+32))
	v1187 = v1175 + base.B2i32(v1183 != int32(0))
	goto L247
L246:
	;
	v1187 = v1175
	goto L247
L247:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+4))
	if v1188 != v1167 {
		v1174 = v1188
		v1175 = v1187
		goto L243
	} else {
		goto L248
	}
L248:
	;
	goto L244
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1159)+uint32(_consts[709]))) = v1196
	goto L251
L250:
	;
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1208 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1208+int32(2816))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L252
	}
L252:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v1220 = v1214
	goto L238
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1220))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_pg_usleep(m, int32(1000000))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_SendPostmasterSignal(m, int32(5))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L255
	}
L255:
	;
	goto L233
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1279 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1283 = F_LWLockAcquire(m, v1279+int32(2816), int32(1))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L257
	}
L257:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, _consts[704]))
	v1288 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v1290 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+20))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+32))
	if v1292 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1393 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1393+int32(2816))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L275
	}
L259:
	;
	v1295 = *(*int64)(unsafe.Add(mBase, uint32(v1292)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1298 = int32(60)
	v1300 = *(*int32)(unsafe.Add(mBase, _consts[710]))
	if v1298 <= v1300 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1303 = v1298
	goto L262
L261:
	;
	v1303 = v1300
	goto L262
L262:
	;
	goto L263
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1314 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1314+int32(2816))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L264
	}
L264:
	;
	if base.B2i32(base.I64_extend_i32_s(v1303*int32(1000))*int64(1000) <= v1275-v1295) == int32(0) {
		goto L233
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	v1324 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1328 = F_LWLockAcquire(m, v1324+int32(2816), int32(0))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L266
	}
L266:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+32))
	if v1332 == int32(0) {
		goto L258
	} else {
		goto L267
	}
L267:
	;
	v1335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1332)+36)) = uint8(v1335)
	v1337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+8)) = v1337
	*(*int64)(unsafe.Add(mBase, uint32(v1332)+24)) = v1337
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+16)) = v1335
	v1344 = v1331 + int32(12)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+16))
	if v1345 == v1335 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+12)) = v1331 + int32(12)
	v1353 = v1344
	goto L270
L269:
	;
	v1353 = v1345
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1332))) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(v1332)+4)) = v1353
	*(*int32)(unsafe.Add(mBase, uint32(v1353))) = v1332
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+16)) = v1332
	v1358 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+32)) = v1358
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1331)+20)) = v1360 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v1368 = F_errstart(m, int32(19), v1358)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L271
	}
L271:
	;
	if v1368 == int32(0) {
		goto L258
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_errmsg(m, int32(447439), int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_errfinish(m, int32(488518), int32(693), int32(274356))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L274
	}
L274:
	;
	goto L258
L275:
	;
	v1398 = v1288 - v1286
	v1399 = int32(0)
	if v1399 < v1398 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1402 = v1398
	goto L278
L277:
	;
	v1402 = v1399
	goto L278
L278:
	;
	if v1291 <= v1402 {
		goto L233
	} else {
		goto L279
	}
L279:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1405 != int32(4088592) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1409 = v1405
	goto L282
L281:
	;
	v1409 = int32(0)
	goto L282
L282:
	;
	if v1409 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v1414 = F_do_start_worker(m)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	v1499 = *(*int64)(unsafe.Add(mBase, uint32(v1496-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	goto L302
L286:
	;
	if v1414 == int32(0) {
		goto L233
	} else {
		goto L287
	}
L287:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1419 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_rebuild_database_list(m, v1414)
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L301
	}
L289:
	;
	if v1419 == int32(4088592) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1429 = v1419
	goto L291
L291:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1429-int32(20))))
	if v1414 == v1439 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L288
L293:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, _consts[710]))
	*(*int64)(unsafe.Add(mBase, uint32(v1429-int32(12)))) = base.I64_extend_i32_s(v1444*int32(1000))*int64(1000) + v1275
	v1453 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1453 == v1429 {
		goto L233
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+4))
	if v1475 != int32(4088592) {
		v1429 = v1475
		goto L291
	} else {
		goto L300
	}
L296:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1429)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1455)+4)) = v1456
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	*(*int32)(unsafe.Add(mBase, uint32(v1456))) = v1458
	v1461 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1461 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1464 = int32(4088592)
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v1464
	v1468 = v1464
	goto L299
L298:
	;
	v1468 = v1461
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = int32(4088592)
	*(*int32)(unsafe.Add(mBase, uint32(v1429)+4)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v1468))) = v1429
	*(*int32)(unsafe.Add(mBase, _consts[697])) = v1429
	goto L233
L300:
	;
	goto L292
L301:
	;
	goto L233
L302:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(0))*int64(1000) <= v1275-v1499) == int32(0) {
		goto L233
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	v1512 = F_do_start_worker(m)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L304
	}
L304:
	;
	if v1512 == int32(0) {
		goto L233
	} else {
		goto L305
	}
L305:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1517 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v629
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v628
	F_rebuild_database_list(m, v1512)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		v1635 = v632
		goto L6
	} else {
		goto L319
	}
L307:
	;
	if v1517 == int32(4088592) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1527 = v1517
	goto L309
L309:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1527-int32(20))))
	if v1512 == v1537 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L306
L311:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, _consts[710]))
	*(*int64)(unsafe.Add(mBase, uint32(v1527-int32(12)))) = base.I64_extend_i32_s(v1542*int32(1000))*int64(1000) + v1275
	v1551 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1551 == v1527 {
		goto L233
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	if v1573 != int32(4088592) {
		v1527 = v1573
		goto L309
	} else {
		goto L318
	}
L314:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1527)))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1553)+4)) = v1554
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1527)))
	*(*int32)(unsafe.Add(mBase, uint32(v1554))) = v1556
	v1559 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	if v1559 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1562 = int32(4088592)
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v1562
	v1566 = v1562
	goto L317
L316:
	;
	v1566 = v1559
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1527))) = int32(4088592)
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+4)) = v1566
	*(*int32)(unsafe.Add(mBase, uint32(v1566))) = v1527
	*(*int32)(unsafe.Add(mBase, _consts[697])) = v1527
	goto L233
L318:
	;
	goto L310
L319:
	;
	goto L233
L320:
	;
	goto L193
L321:
	;
	goto L5
L322:
	;
	v1645 = int32(v1641)
	m.G0 = v1635
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+4))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1645)))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1648)))
	if v21+int32(36) == v1652 {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	m.ExcPending = 1
	goto L331
L324:
	;
	if v1655 != 0 {
		goto L328
	} else {
		goto L329
	}
L325:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+4))
	v1655 = v1654
	goto L327
L326:
	;
	v1655 = int32(0)
	goto L327
L327:
	;
	goto L324
L328:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v22 = v1657
	v23 = v1656
	v24 = v1655
	v25 = v1647
	v27 = v1635
	goto L1
L329:
	;
	goto L330
L330:
	;
	F___wasm_longjmp(m, v1648, v1647)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	return
L332:
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[120]))
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
					*(*int32)(unsafe.Add(mBase, _consts[148])) = int32(4453796)
					*(*int32)(unsafe.Add(mBase, _consts[157])) = int32(-1)
					v21 = int32(4393996)
					v22 = *(*int32)(unsafe.Add(mBase, _consts[120]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[120])) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v22+int32(20))+12)) = v24
					v31 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(1)
					if v32 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
						F_s_lock(m, v36, int32(491328), int32(1071), int32(299371))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v22)+52)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
							v47 = *(*int32)(unsafe.Add(mBase, _consts[156]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
							v50 = *(*int32)(unsafe.Add(mBase, _consts[414]))
							v55 = base.I32_div_s(v50+v48*int32(15), int32(16))
							v57 = *(*int32)(unsafe.Add(mBase, _consts[156]))
							*(*int32)(unsafe.Add(mBase, uint32(v57)+68)) = v55
							v60 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
							return
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v22)+52)) = int64(4294967295)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(0)
						v47 = *(*int32)(unsafe.Add(mBase, _consts[156]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
						v50 = *(*int32)(unsafe.Add(mBase, _consts[414]))
						v55 = base.I32_div_s(v50+v48*int32(15), int32(16))
						v57 = *(*int32)(unsafe.Add(mBase, _consts[156]))
						*(*int32)(unsafe.Add(mBase, uint32(v57)+68)) = v55
						v60 = *(*int32)(unsafe.Add(mBase, _consts[1120]))
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
			F_errmsg_internal(m, int32(127726), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				F_errfinish(m, int32(491328), int32(1050), int32(299371))
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
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
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
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v213 int64
	_ = v213
	var v220 int64
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v242 int64
	_ = v242
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v256 int32
	_ = v256
	var v271 int64
	_ = v271
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v290 int64
	_ = v290
	var v293 int64
	_ = v293
	var v295 int64
	_ = v295
	var v298 int64
	_ = v298
	var v305 int64
	_ = v305
	var v309 int64
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v324 int32
	_ = v324
	var v335 int64
	_ = v335
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v355 int64
	_ = v355
	var v360 int64
	_ = v360
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v401 int64
	_ = v401
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v415 int32
	_ = v415
	var v430 int64
	_ = v430
	var v434 int64
	_ = v434
	var v435 int64
	_ = v435
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v447 int64
	_ = v447
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v464 int64
	_ = v464
	var v467 int32
	_ = v467
	var v473 int64
	_ = v473
	var v476 int64
	_ = v476
	var v479 int64
	_ = v479
	var v485 int64
	_ = v485
	var v486 int64
	_ = v486
	var v490 int64
	_ = v490
	var v491 int64
	_ = v491
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
			if v23 == v19 {
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
			v108 = int32(32767)
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
				v127 = base.I32_wrap_i64(base.I64_clz(v121) + base.I64_extend_i32_u(v120<<(uint(int32(6))%32)))
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
				v155 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
				v156 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
				v157 = v156
				v158 = v155
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
				v205 = v109
				v206 = v162
			} else {
				v164 = v16 + int32(80)
				v166 = base.B2i32(v162 == int64(0))
				if v162 == int64(0) {
					v167 = v160
				} else {
					v167 = v162
				}
				v173 = base.I32_wrap_i64(base.I64_clz(v167) + base.I64_extend_i32_u(v166<<(uint(int32(6))%32)))
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
				v205 = int32(16) - v173
				v206 = v202
			}
			v207 = int64(3)
			v209 = int64(61)
			v213 = v206<<(uint(v207)%64) | int64(base.Ui64(v203)>>(uint(v209)%64)) | int64(2251799813685248)
			v220 = v203 << (uint(v207) % 64)
			if v159 == v205 {
				v289 = v213
				v290 = v220
			} else {
				v223 = v159 - v205
				if base.Ui32(int32(127)) < base.Ui32(v223) {
					v289 = int64(0)
					v290 = int64(1)
				} else {
					v229 = v16 - int32(-64)
					v231 = int32(128) - v223
					if v231&int32(64) != 0 {
						v250 = int64(0)
						v251 = v220 << (uint(base.I64_extend_i32_u(v231+int32(-64))) % 64)
					} else {
						if v231 == int32(0) {
							v250 = v220
							v251 = v213
						} else {
							v242 = base.I64_extend_i32_u(v231)
							v250 = v220 << (uint(v242) % 64)
							v251 = v213<<(uint(v242)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v231))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v229))) = v250
					*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v251
					v256 = v16 + int32(48)
					if v223&int32(64) != 0 {
						v275 = int64(base.Ui64(v213) >> (uint(base.I64_extend_i32_u(v223+int32(-64))) % 64))
						v276 = int64(0)
					} else {
						if v223 == int32(0) {
							v275 = v220
							v276 = v213
						} else {
							v271 = base.I64_extend_i32_u(v223)
							v275 = v213<<(uint(base.I64_extend_i32_u(int32(64)-v223))%64) | int64(base.Ui64(v220)>>(uint(v271)%64))
							v276 = int64(base.Ui64(v213) >> (uint(v271) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v256))) = v275
					*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v276
					v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
					v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
					v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
					v288 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
					v289 = v288
					v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
				}
			}
			v293 = v157<<(uint(v207)%64) | int64(base.Ui64(v158)>>(uint(v209)%64)) | int64(2251799813685248)
			v295 = v158 << (uint(int64(3)) % 64)
			if l2^l4 < int64(0) {
				v298 = int64(0)
				if v290^v295|(v289^v293) == v298 {
					v490 = v298
					v491 = v298
				} else {
					v305 = v295 - v290
					v309 = v293 - v289 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v290)))
					if base.Ui64(int64(2251799813685247)) < base.Ui64(v309) {
						v372 = v305
						v374 = v309
						v375 = v159
					} else {
						v313 = v16 + int32(32)
						v315 = base.B2i32(v309 == int64(0))
						if v309 == int64(0) {
							v316 = v305
						} else {
							v316 = v309
						}
						v324 = base.I32_wrap_i64(base.I64_clz(v316)+base.I64_extend_i32_u(v315<<(uint(int32(6))%32))) - int32(12)
						if v324&int32(64) != 0 {
							v343 = int64(0)
							v344 = v305 << (uint(base.I64_extend_i32_u(v324+int32(-64))) % 64)
						} else {
							if v324 == int32(0) {
								v343 = v305
								v344 = v309
							} else {
								v335 = base.I64_extend_i32_u(v324)
								v343 = v305 << (uint(v335) % 64)
								v344 = v309<<(uint(v335)%64) | int64(base.Ui64(v305)>>(uint(base.I64_extend_i32_u(int32(64)-v324))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v313))) = v343
						*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = v344
						v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
						v350 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
						v372 = v350
						v374 = v349
						v375 = v159 - v324
					}
					v378 = v101 & int64(-9223372036854775807-1)
					if int32(32767) <= v375 {
						v490 = int64(0)
						v491 = v378 | int64(9223090561878065152)
					} else {
						v384 = int32(0)
						if v384 < v375 {
							v448 = v372
							v449 = v374
							v450 = v375
						} else {
							v388 = v16 + int32(16)
							v390 = v375 + int32(127)
							if v390&int32(64) != 0 {
								v409 = int64(0)
								v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
							} else {
								if v390 == int32(0) {
									v409 = v372
									v410 = v374
								} else {
									v401 = base.I64_extend_i32_u(v390)
									v409 = v372 << (uint(v401) % 64)
									v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
							*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
							v415 = int32(1) - v375
							if v415&int32(64) != 0 {
								v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
								v435 = int64(0)
							} else {
								if v415 == int32(0) {
									v434 = v372
									v435 = v374
								} else {
									v430 = base.I64_extend_i32_u(v415)
									v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
									v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
								}
							}
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
							v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
							v449 = v447
							v450 = v384
						}
						v453 = int64(3)
						v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
						v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
						v467 = base.I32_wrap_i64(v448) & int32(7)
						if v467 != int32(4) {
							v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
							v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
							if v467 == int32(0) {
								v490 = v473
								v491 = v476
							} else {
								v485 = v473
								v486 = v476
								v490 = v485
								v491 = v486
							}
						} else {
							v479 = v455 + v455&int64(1)
							v485 = v479
							v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
							v490 = v485
							v491 = v486
						}
					}
				}
			} else {
				v351 = v290 + v295
				v355 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v351) < base.Ui64(v290))) + (v289 + v293)
				if v355&int64(4503599627370496) == int64(0) {
					v372 = v351
					v374 = v355
					v375 = v159
				} else {
					v360 = int64(1)
					v372 = v290&v360 | (v355<<(uint(int64(63))%64) | int64(base.Ui64(v351)>>(uint(v360)%64)))
					v374 = int64(base.Ui64(v355) >> (uint(v360) % 64))
					v375 = v159 + int32(1)
				}
				v378 = v101 & int64(-9223372036854775807-1)
				if int32(32767) <= v375 {
					v490 = int64(0)
					v491 = v378 | int64(9223090561878065152)
				} else {
					v384 = int32(0)
					if v384 < v375 {
						v448 = v372
						v449 = v374
						v450 = v375
					} else {
						v388 = v16 + int32(16)
						v390 = v375 + int32(127)
						if v390&int32(64) != 0 {
							v409 = int64(0)
							v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
						} else {
							if v390 == int32(0) {
								v409 = v372
								v410 = v374
							} else {
								v401 = base.I64_extend_i32_u(v390)
								v409 = v372 << (uint(v401) % 64)
								v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
						*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
						v415 = int32(1) - v375
						if v415&int32(64) != 0 {
							v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
							v435 = int64(0)
						} else {
							if v415 == int32(0) {
								v434 = v372
								v435 = v374
							} else {
								v430 = base.I64_extend_i32_u(v415)
								v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
								v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
						v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
						v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
						v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
						v449 = v447
						v450 = v384
					}
					v453 = int64(3)
					v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
					v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
					v467 = base.I32_wrap_i64(v448) & int32(7)
					if v467 != int32(4) {
						v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
						v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
						if v467 == int32(0) {
							v490 = v473
							v491 = v476
						} else {
							v485 = v473
							v486 = v476
							v490 = v485
							v491 = v486
						}
					} else {
						v479 = v455 + v455&int64(1)
						v485 = v479
						v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
						v490 = v485
						v491 = v486
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
				v490 = l1
				v491 = l2 | int64(140737488355328)
			} else {
				v54 = int64(9223090561878065152)
				if v19 == v54 {
					v58 = base.B2i32(l3 == int64(0))
				} else {
					v58 = base.B2i32(base.Ui64(v19) < base.Ui64(v54))
				}
				if v58 == int32(0) {
					v490 = l3
					v491 = l4 | int64(140737488355328)
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
						v490 = v78
						v491 = v76
					} else {
						if l3|(v19^int64(9223090561878065152)) == int64(0) {
							v490 = l3
							v491 = l4
						} else {
							if l1|v23 == int64(0) {
								if l3|v19 != int64(0) {
									v490 = l3
									v491 = l4
								} else {
									v490 = l1 & l3
									v491 = l2 & l4
								}
							} else {
								if l3|v19 != int64(0) {
									if v23 == v19 {
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
									v108 = int32(32767)
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
										v127 = base.I32_wrap_i64(base.I64_clz(v121) + base.I64_extend_i32_u(v120<<(uint(int32(6))%32)))
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
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
										v156 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
										v157 = v156
										v158 = v155
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
										v205 = v109
										v206 = v162
									} else {
										v164 = v16 + int32(80)
										v166 = base.B2i32(v162 == int64(0))
										if v162 == int64(0) {
											v167 = v160
										} else {
											v167 = v162
										}
										v173 = base.I32_wrap_i64(base.I64_clz(v167) + base.I64_extend_i32_u(v166<<(uint(int32(6))%32)))
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
										v205 = int32(16) - v173
										v206 = v202
									}
									v207 = int64(3)
									v209 = int64(61)
									v213 = v206<<(uint(v207)%64) | int64(base.Ui64(v203)>>(uint(v209)%64)) | int64(2251799813685248)
									v220 = v203 << (uint(v207) % 64)
									if v159 == v205 {
										v289 = v213
										v290 = v220
									} else {
										v223 = v159 - v205
										if base.Ui32(int32(127)) < base.Ui32(v223) {
											v289 = int64(0)
											v290 = int64(1)
										} else {
											v229 = v16 - int32(-64)
											v231 = int32(128) - v223
											if v231&int32(64) != 0 {
												v250 = int64(0)
												v251 = v220 << (uint(base.I64_extend_i32_u(v231+int32(-64))) % 64)
											} else {
												if v231 == int32(0) {
													v250 = v220
													v251 = v213
												} else {
													v242 = base.I64_extend_i32_u(v231)
													v250 = v220 << (uint(v242) % 64)
													v251 = v213<<(uint(v242)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v231))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v229))) = v250
											*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v251
											v256 = v16 + int32(48)
											if v223&int32(64) != 0 {
												v275 = int64(base.Ui64(v213) >> (uint(base.I64_extend_i32_u(v223+int32(-64))) % 64))
												v276 = int64(0)
											} else {
												if v223 == int32(0) {
													v275 = v220
													v276 = v213
												} else {
													v271 = base.I64_extend_i32_u(v223)
													v275 = v213<<(uint(base.I64_extend_i32_u(int32(64)-v223))%64) | int64(base.Ui64(v220)>>(uint(v271)%64))
													v276 = int64(base.Ui64(v213) >> (uint(v271) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v256))) = v275
											*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v276
											v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
											v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
											v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
											v288 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
											v289 = v288
											v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
										}
									}
									v293 = v157<<(uint(v207)%64) | int64(base.Ui64(v158)>>(uint(v209)%64)) | int64(2251799813685248)
									v295 = v158 << (uint(int64(3)) % 64)
									if l2^l4 < int64(0) {
										v298 = int64(0)
										if v290^v295|(v289^v293) == v298 {
											v490 = v298
											v491 = v298
										} else {
											v305 = v295 - v290
											v309 = v293 - v289 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v290)))
											if base.Ui64(int64(2251799813685247)) < base.Ui64(v309) {
												v372 = v305
												v374 = v309
												v375 = v159
											} else {
												v313 = v16 + int32(32)
												v315 = base.B2i32(v309 == int64(0))
												if v309 == int64(0) {
													v316 = v305
												} else {
													v316 = v309
												}
												v324 = base.I32_wrap_i64(base.I64_clz(v316)+base.I64_extend_i32_u(v315<<(uint(int32(6))%32))) - int32(12)
												if v324&int32(64) != 0 {
													v343 = int64(0)
													v344 = v305 << (uint(base.I64_extend_i32_u(v324+int32(-64))) % 64)
												} else {
													if v324 == int32(0) {
														v343 = v305
														v344 = v309
													} else {
														v335 = base.I64_extend_i32_u(v324)
														v343 = v305 << (uint(v335) % 64)
														v344 = v309<<(uint(v335)%64) | int64(base.Ui64(v305)>>(uint(base.I64_extend_i32_u(int32(64)-v324))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v313))) = v343
												*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = v344
												v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
												v350 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
												v372 = v350
												v374 = v349
												v375 = v159 - v324
											}
											v378 = v101 & int64(-9223372036854775807-1)
											if int32(32767) <= v375 {
												v490 = int64(0)
												v491 = v378 | int64(9223090561878065152)
											} else {
												v384 = int32(0)
												if v384 < v375 {
													v448 = v372
													v449 = v374
													v450 = v375
												} else {
													v388 = v16 + int32(16)
													v390 = v375 + int32(127)
													if v390&int32(64) != 0 {
														v409 = int64(0)
														v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
													} else {
														if v390 == int32(0) {
															v409 = v372
															v410 = v374
														} else {
															v401 = base.I64_extend_i32_u(v390)
															v409 = v372 << (uint(v401) % 64)
															v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
														}
													}
													*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
													*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
													v415 = int32(1) - v375
													if v415&int32(64) != 0 {
														v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
														v435 = int64(0)
													} else {
														if v415 == int32(0) {
															v434 = v372
															v435 = v374
														} else {
															v430 = base.I64_extend_i32_u(v415)
															v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
															v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
														}
													}
													*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
													*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
													v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
													v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
													v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
													v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
													v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
													v449 = v447
													v450 = v384
												}
												v453 = int64(3)
												v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
												v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
												v467 = base.I32_wrap_i64(v448) & int32(7)
												if v467 != int32(4) {
													v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
													v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
													if v467 == int32(0) {
														v490 = v473
														v491 = v476
													} else {
														v485 = v473
														v486 = v476
														v490 = v485
														v491 = v486
													}
												} else {
													v479 = v455 + v455&int64(1)
													v485 = v479
													v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
													v490 = v485
													v491 = v486
												}
											}
										}
									} else {
										v351 = v290 + v295
										v355 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v351) < base.Ui64(v290))) + (v289 + v293)
										if v355&int64(4503599627370496) == int64(0) {
											v372 = v351
											v374 = v355
											v375 = v159
										} else {
											v360 = int64(1)
											v372 = v290&v360 | (v355<<(uint(int64(63))%64) | int64(base.Ui64(v351)>>(uint(v360)%64)))
											v374 = int64(base.Ui64(v355) >> (uint(v360) % 64))
											v375 = v159 + int32(1)
										}
										v378 = v101 & int64(-9223372036854775807-1)
										if int32(32767) <= v375 {
											v490 = int64(0)
											v491 = v378 | int64(9223090561878065152)
										} else {
											v384 = int32(0)
											if v384 < v375 {
												v448 = v372
												v449 = v374
												v450 = v375
											} else {
												v388 = v16 + int32(16)
												v390 = v375 + int32(127)
												if v390&int32(64) != 0 {
													v409 = int64(0)
													v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
												} else {
													if v390 == int32(0) {
														v409 = v372
														v410 = v374
													} else {
														v401 = base.I64_extend_i32_u(v390)
														v409 = v372 << (uint(v401) % 64)
														v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
												*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
												v415 = int32(1) - v375
												if v415&int32(64) != 0 {
													v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
													v435 = int64(0)
												} else {
													if v415 == int32(0) {
														v434 = v372
														v435 = v374
													} else {
														v430 = base.I64_extend_i32_u(v415)
														v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
														v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
												v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
												v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
												v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
												v449 = v447
												v450 = v384
											}
											v453 = int64(3)
											v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
											v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
											v467 = base.I32_wrap_i64(v448) & int32(7)
											if v467 != int32(4) {
												v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
												v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
												if v467 == int32(0) {
													v490 = v473
													v491 = v476
												} else {
													v485 = v473
													v486 = v476
													v490 = v485
													v491 = v486
												}
											} else {
												v479 = v455 + v455&int64(1)
												v485 = v479
												v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
												v490 = v485
												v491 = v486
											}
										}
									}
								} else {
									v490 = l1
									v491 = l2
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
			v490 = l1
			v491 = l2 | int64(140737488355328)
		} else {
			v54 = int64(9223090561878065152)
			if v19 == v54 {
				v58 = base.B2i32(l3 == int64(0))
			} else {
				v58 = base.B2i32(base.Ui64(v19) < base.Ui64(v54))
			}
			if v58 == int32(0) {
				v490 = l3
				v491 = l4 | int64(140737488355328)
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
					v490 = v78
					v491 = v76
				} else {
					if l3|(v19^int64(9223090561878065152)) == int64(0) {
						v490 = l3
						v491 = l4
					} else {
						if l1|v23 == int64(0) {
							if l3|v19 != int64(0) {
								v490 = l3
								v491 = l4
							} else {
								v490 = l1 & l3
								v491 = l2 & l4
							}
						} else {
							if l3|v19 != int64(0) {
								if v23 == v19 {
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
								v108 = int32(32767)
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
									v127 = base.I32_wrap_i64(base.I64_clz(v121) + base.I64_extend_i32_u(v120<<(uint(int32(6))%32)))
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
									v155 = *(*int64)(unsafe.Add(mBase, uint32(v16)+96))
									v156 = *(*int64)(unsafe.Add(mBase, uint32(v16)+104))
									v157 = v156
									v158 = v155
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
									v205 = v109
									v206 = v162
								} else {
									v164 = v16 + int32(80)
									v166 = base.B2i32(v162 == int64(0))
									if v162 == int64(0) {
										v167 = v160
									} else {
										v167 = v162
									}
									v173 = base.I32_wrap_i64(base.I64_clz(v167) + base.I64_extend_i32_u(v166<<(uint(int32(6))%32)))
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
									v205 = int32(16) - v173
									v206 = v202
								}
								v207 = int64(3)
								v209 = int64(61)
								v213 = v206<<(uint(v207)%64) | int64(base.Ui64(v203)>>(uint(v209)%64)) | int64(2251799813685248)
								v220 = v203 << (uint(v207) % 64)
								if v159 == v205 {
									v289 = v213
									v290 = v220
								} else {
									v223 = v159 - v205
									if base.Ui32(int32(127)) < base.Ui32(v223) {
										v289 = int64(0)
										v290 = int64(1)
									} else {
										v229 = v16 - int32(-64)
										v231 = int32(128) - v223
										if v231&int32(64) != 0 {
											v250 = int64(0)
											v251 = v220 << (uint(base.I64_extend_i32_u(v231+int32(-64))) % 64)
										} else {
											if v231 == int32(0) {
												v250 = v220
												v251 = v213
											} else {
												v242 = base.I64_extend_i32_u(v231)
												v250 = v220 << (uint(v242) % 64)
												v251 = v213<<(uint(v242)%64) | int64(base.Ui64(v220)>>(uint(base.I64_extend_i32_u(int32(64)-v231))%64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v229))) = v250
										*(*int64)(unsafe.Add(mBase, uint32(v229)+8)) = v251
										v256 = v16 + int32(48)
										if v223&int32(64) != 0 {
											v275 = int64(base.Ui64(v213) >> (uint(base.I64_extend_i32_u(v223+int32(-64))) % 64))
											v276 = int64(0)
										} else {
											if v223 == int32(0) {
												v275 = v220
												v276 = v213
											} else {
												v271 = base.I64_extend_i32_u(v223)
												v275 = v213<<(uint(base.I64_extend_i32_u(int32(64)-v223))%64) | int64(base.Ui64(v220)>>(uint(v271)%64))
												v276 = int64(base.Ui64(v213) >> (uint(v271) % 64))
											}
										}
										*(*int64)(unsafe.Add(mBase, uint32(v256))) = v275
										*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v276
										v280 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
										v281 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
										v282 = *(*int64)(unsafe.Add(mBase, uint32(v16)+72))
										v288 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
										v289 = v288
										v290 = v280 | base.I64_extend_i32_u(base.B2i32(v281|v282 != int64(0)))
									}
								}
								v293 = v157<<(uint(v207)%64) | int64(base.Ui64(v158)>>(uint(v209)%64)) | int64(2251799813685248)
								v295 = v158 << (uint(int64(3)) % 64)
								if l2^l4 < int64(0) {
									v298 = int64(0)
									if v290^v295|(v289^v293) == v298 {
										v490 = v298
										v491 = v298
									} else {
										v305 = v295 - v290
										v309 = v293 - v289 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v295) < base.Ui64(v290)))
										if base.Ui64(int64(2251799813685247)) < base.Ui64(v309) {
											v372 = v305
											v374 = v309
											v375 = v159
										} else {
											v313 = v16 + int32(32)
											v315 = base.B2i32(v309 == int64(0))
											if v309 == int64(0) {
												v316 = v305
											} else {
												v316 = v309
											}
											v324 = base.I32_wrap_i64(base.I64_clz(v316)+base.I64_extend_i32_u(v315<<(uint(int32(6))%32))) - int32(12)
											if v324&int32(64) != 0 {
												v343 = int64(0)
												v344 = v305 << (uint(base.I64_extend_i32_u(v324+int32(-64))) % 64)
											} else {
												if v324 == int32(0) {
													v343 = v305
													v344 = v309
												} else {
													v335 = base.I64_extend_i32_u(v324)
													v343 = v305 << (uint(v335) % 64)
													v344 = v309<<(uint(v335)%64) | int64(base.Ui64(v305)>>(uint(base.I64_extend_i32_u(int32(64)-v324))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v313))) = v343
											*(*int64)(unsafe.Add(mBase, uint32(v313)+8)) = v344
											v349 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
											v350 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
											v372 = v350
											v374 = v349
											v375 = v159 - v324
										}
										v378 = v101 & int64(-9223372036854775807-1)
										if int32(32767) <= v375 {
											v490 = int64(0)
											v491 = v378 | int64(9223090561878065152)
										} else {
											v384 = int32(0)
											if v384 < v375 {
												v448 = v372
												v449 = v374
												v450 = v375
											} else {
												v388 = v16 + int32(16)
												v390 = v375 + int32(127)
												if v390&int32(64) != 0 {
													v409 = int64(0)
													v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
												} else {
													if v390 == int32(0) {
														v409 = v372
														v410 = v374
													} else {
														v401 = base.I64_extend_i32_u(v390)
														v409 = v372 << (uint(v401) % 64)
														v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
												*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
												v415 = int32(1) - v375
												if v415&int32(64) != 0 {
													v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
													v435 = int64(0)
												} else {
													if v415 == int32(0) {
														v434 = v372
														v435 = v374
													} else {
														v430 = base.I64_extend_i32_u(v415)
														v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
														v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
													}
												}
												*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
												*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
												v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
												v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
												v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
												v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
												v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
												v449 = v447
												v450 = v384
											}
											v453 = int64(3)
											v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
											v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
											v467 = base.I32_wrap_i64(v448) & int32(7)
											if v467 != int32(4) {
												v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
												v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
												if v467 == int32(0) {
													v490 = v473
													v491 = v476
												} else {
													v485 = v473
													v486 = v476
													v490 = v485
													v491 = v486
												}
											} else {
												v479 = v455 + v455&int64(1)
												v485 = v479
												v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
												v490 = v485
												v491 = v486
											}
										}
									}
								} else {
									v351 = v290 + v295
									v355 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v351) < base.Ui64(v290))) + (v289 + v293)
									if v355&int64(4503599627370496) == int64(0) {
										v372 = v351
										v374 = v355
										v375 = v159
									} else {
										v360 = int64(1)
										v372 = v290&v360 | (v355<<(uint(int64(63))%64) | int64(base.Ui64(v351)>>(uint(v360)%64)))
										v374 = int64(base.Ui64(v355) >> (uint(v360) % 64))
										v375 = v159 + int32(1)
									}
									v378 = v101 & int64(-9223372036854775807-1)
									if int32(32767) <= v375 {
										v490 = int64(0)
										v491 = v378 | int64(9223090561878065152)
									} else {
										v384 = int32(0)
										if v384 < v375 {
											v448 = v372
											v449 = v374
											v450 = v375
										} else {
											v388 = v16 + int32(16)
											v390 = v375 + int32(127)
											if v390&int32(64) != 0 {
												v409 = int64(0)
												v410 = v372 << (uint(base.I64_extend_i32_u(v375+int32(63))) % 64)
											} else {
												if v390 == int32(0) {
													v409 = v372
													v410 = v374
												} else {
													v401 = base.I64_extend_i32_u(v390)
													v409 = v372 << (uint(v401) % 64)
													v410 = v374<<(uint(v401)%64) | int64(base.Ui64(v372)>>(uint(base.I64_extend_i32_u(int32(64)-v390))%64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v388))) = v409
											*(*int64)(unsafe.Add(mBase, uint32(v388)+8)) = v410
											v415 = int32(1) - v375
											if v415&int32(64) != 0 {
												v434 = int64(base.Ui64(v374) >> (uint(base.I64_extend_i32_u(v415+int32(-64))) % 64))
												v435 = int64(0)
											} else {
												if v415 == int32(0) {
													v434 = v372
													v435 = v374
												} else {
													v430 = base.I64_extend_i32_u(v415)
													v434 = v374<<(uint(base.I64_extend_i32_u(int32(64)-v415))%64) | int64(base.Ui64(v372)>>(uint(v430)%64))
													v435 = int64(base.Ui64(v374) >> (uint(v430) % 64))
												}
											}
											*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
											*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v435
											v439 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
											v440 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
											v441 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
											v447 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
											v448 = v439 | base.I64_extend_i32_u(base.B2i32(v440|v441 != int64(0)))
											v449 = v447
											v450 = v384
										}
										v453 = int64(3)
										v455 = v449<<(uint(int64(61))%64) | int64(base.Ui64(v448)>>(uint(v453)%64))
										v464 = int64(base.Ui64(v449)>>(uint(v453)%64))&int64(281474976710655) | base.I64_extend_i32_u(v450)<<(uint(int64(48))%64) | v378
										v467 = base.I32_wrap_i64(v448) & int32(7)
										if v467 != int32(4) {
											v473 = v455 + base.I64_extend_i32_u(base.B2i32(base.Ui32(int32(4)) < base.Ui32(v467)))
											v476 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v473) < base.Ui64(v455)))
											if v467 == int32(0) {
												v490 = v473
												v491 = v476
											} else {
												v485 = v473
												v486 = v476
												v490 = v485
												v491 = v486
											}
										} else {
											v479 = v455 + v455&int64(1)
											v485 = v479
											v486 = v464 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v479) < base.Ui64(v455)))
											v490 = v485
											v491 = v486
										}
									}
								}
							} else {
								v490 = l1
								v491 = l2
							}
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v490
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v491
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
	var v63 int32
	_ = v63
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
	var v147 int64
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v216 int32
	_ = v216
	var v223 int64
	_ = v223
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L84
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
	return v223
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
		v223 = v21
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
		v63 = v32
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = v63
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
		v63 = v56
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v63 = v56
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
		v223 = v21
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
	if l3 != v112 {
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
		v223 = v112
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
	if v143 == int32(0) {
		v213 = v134
		v214 = v136
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v223 = v213
	goto L9
L54:
	;
	v216 = v129 + int32(1)
	if v216 != v75 {
		v129 = v216
		v134 = v213
		v136 = v214
		goto L52
	} else {
		goto L83
	}
L55:
	;
	if l1 == v143 {
		v213 = v134
		v214 = v136
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	if v147&v136 == int64(0) {
		v213 = v134
		v214 = v136
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v151 = F_superuser_arg(m, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v151 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = int32(0)
	v158 = F_roles_is_member_of(m, l1, int32(1), v156, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	v203 = v201&l3 | v134
	if l4 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L62:
	;
	v160 = int32(0)
	if v158 == v160 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v198 == int32(0) {
		v213 = v134
		v214 = v136
		goto L54
	} else {
		goto L76
	}
L64:
	;
	v198 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v166 <= int32(0) {
		v191 = v160
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v198 = v191
	goto L63
L68:
	;
	v169 = int32(0)
	if v169 < v166 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v172 = v166
	goto L71
L70:
	;
	v172 = v169
	goto L71
L71:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v175 = int32(0)
	goto L72
L72:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v173+v175<<(uint(int32(2))%32))))
	v184 = base.B2i32(v183 == v143)
	if v183 == v143 {
		v191 = v184
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v191 = v184
	goto L67
L74:
	;
	v186 = v175 + int32(1)
	if v186 != v172 {
		v175 = v186
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	goto L61
L77:
	;
	v213 = v203
	v214 = l3 & (v203 ^ int64(-1))
	goto L54
L78:
	;
	if l3 != v203 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v203 != int64(0) {
		v223 = v203
		goto L9
	} else {
		goto L82
	}
L81:
	;
	return l3
L82:
	;
	goto L77
L83:
	;
	goto L53
L84:
	;
	F_errmsg_internal(m, int32(524738), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(488995), int32(1403), int32(309798))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
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
			F_errmsg(m, int32(436893), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(488995), int32(1607), int32(336602))
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
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
	v88 = F_lappend(m, v84, l2)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L30
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
	v84 = int32(0)
	goto L2
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v16 {
		v84 = v18
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
		v60 = v16
		v62 = v18
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v62 != 0 {
		v16 = v60 + int32(1)
		v18 = v62
		goto L6
	} else {
		goto L29
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
	if v44 == int32(-3) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v35 != v39 {
		v44 = v38
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 == v42 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	if v38 != v32 {
		v44 = v38
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L1
L18:
	;
	v44 = v42
	goto L12
L19:
	;
	v54 = F_list_delete_nth_cell(m, v18, v16)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v47 != int32(-3) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v44 != v32 {
		v60 = v16
		v62 = v18
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v44 != v32 {
		v60 = v16
		v62 = v18
		goto L9
	} else {
		goto L26
	}
L24:
	;
	if v35 == v47 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v60 = v16
	v62 = v18
	goto L9
L26:
	;
	goto L19
L27:
	;
	return
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v54
	v60 = v16 - int32(1)
	v62 = v54
	goto L9
L29:
	;
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+36))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v93 == v94 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v96 | int32(2)
	return
L32:
	;
	goto L33
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v102 = F_pg_reg_getnumoutarcs(m, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v106 = F_palloc(m, v102<<(uint(int32(3))%32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_pg_reg_getoutarcs(m, v108, v109, v106, v102)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	if int32(0) < v102 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v121 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_pfree(m, v106)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L27
	} else {
		goto L76
	}
L40:
	;
	v125 = int32(-4)
	v129 = v106 + v121<<(uint(int32(3))%32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
	v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+40)))
	if v130 != v133 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L39
L42:
	;
	v222 = v121 + int32(1)
	if v222 != v102 {
		v121 = v222
		goto L40
	} else {
		goto L75
	}
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v208 = F_palloc(m, int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L27
	} else {
		goto L73
	}
L44:
	;
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v132)+42)))
	v138 = base.B2i32(v130 == v135)
	goto L46
L45:
	;
	v138 = int32(1)
	goto L46
L46:
	;
	if v138 != 0 {
		v204 = v125
		v205 = v125
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v139 = int32(-3)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+24))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+44)))
	if v140 != v143 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+46)))
	v148 = base.B2i32(v140 == v145)
	goto L50
L49:
	;
	v148 = int32(1)
	goto L50
L50:
	;
	if v148 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v204 = v139
	v205 = int32(-3)
	goto L43
L52:
	;
	goto L53
L53:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v150 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v204 = v139
	v205 = int32(-3)
	goto L43
L55:
	;
	goto L56
L56:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = v155 + v150*int32(12)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != int32(1) {
		v204 = v139
		v205 = int32(-3)
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	if v162 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v187 <= int32(0) {
		goto L42
	} else {
		goto L67
	}
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v165 + int32(4) {
	case 0:
		goto L62
	case 1:
		goto L60
	default:
		goto L61
	}
L60:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v176 = F_palloc(m, int32(12))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L27
	} else {
		goto L65
	}
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v171 != int32(-4) {
		goto L58
	} else {
		goto L64
	}
L62:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v168 == int32(-4) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L58
L64:
	;
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v174
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = int64(-12884901892)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v182 = F_lappend(m, v181, v176)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L27
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v182
	goto L58
L67:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v192 + int32(4) {
	case 0:
		goto L69
	case 1:
		v204 = v190
		v205 = v191
		goto L43
	default:
		goto L68
	}
L68:
	;
	v200 = int32(-4)
	if v190 != v200 {
		goto L42
	} else {
		goto L72
	}
L69:
	;
	if v191 != int32(-4) {
		goto L42
	} else {
		goto L70
	}
L70:
	;
	v197 = int32(-4)
	if v190 == v197 {
		v204 = v197
		v205 = v191
		goto L43
	} else {
		goto L71
	}
L71:
	;
	goto L42
L72:
	;
	v204 = v200
	v205 = v191
	goto L43
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v204
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v214 = F_lappend(m, v213, v208)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L27
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v214
	goto L42
L75:
	;
	goto L41
L76:
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
		v19 = int32(base.Ui32(v11+int32(262120)) >> (uint(int32(2)) % 32))
	} else {
		v19 = int32(0)
	}
	if base.Ui32(l3) <= base.Ui32(v19&int32(65535)) {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32)+l0)+20))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+v26&int32(32767))))
		v31 = int32(3)
		if v30&v31 != v31 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(223807), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errfinish(m, int32(489526), int32(58), int32(378598))
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
							F_errmsg_internal(m, int32(400754), v9)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(489526), int32(70), int32(378598))
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
					F_errmsg_internal(m, int32(400754), v9)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(489526), int32(70), int32(378598))
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = int32(-1)
	v27 = v23 + (v24 ^ v25)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = v28 + (v29 ^ v25)
	if v32 < v27 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = v27
	goto L3
L2:
	;
	v34 = v32
	goto L3
L3:
	;
	v36 = v34 + int32(1)
	if v29 < v24 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = v24
	goto L6
L5:
	;
	v38 = v29
	goto L6
L6:
	;
	v39 = int32(1)
	v40 = v38 + v39
	v41 = v36 + v40
	if v41 <= v39 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = int32(1)
	goto L9
L8:
	;
	v44 = v41
	goto L9
L9:
	;
	v46 = v44 << (uint(int32(1)) % 32)
	v49 = F_palloc(m, v46+int32(2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v51 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = v36 + v53
	v63 = v51
	v64 = v36 + v55
	v69 = v44
	goto L12
L12:
	;
	v76 = v64 - int32(1)
	if v76 < int32(0) {
		v85 = v63
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v18 < v19 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v86 = int32(1)
	v92 = v61 - v86
	if v92 < int32(0) {
		v101 = v85
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if v23 <= v76 {
		v85 = v63
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v76<<(uint(int32(1))%32)))))
	v85 = v63 + v83
	goto L14
L17:
	;
	v105 = base.B2i32(int32(9999) < v101)
	if int32(9999) < v101 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v28 <= v92 {
		v101 = v85
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20+v92<<(uint(int32(1))%32)))))
	v101 = v85 + v99
	goto L17
L20:
	;
	v106 = v101 - int32(10000)
	goto L22
L21:
	;
	v106 = v101
	goto L22
L22:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v69<<(uint(v86)%32)))) = uint16(v106)
	if base.Ui32(int32(1)) < base.Ui32(v69) {
		v61 = v92
		v63 = v105
		v64 = v76
		v69 = v69 - v86
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L13
L24:
	;
	v111 = v19
	goto L26
L25:
	;
	v111 = v18
	goto L26
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v112 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_pfree(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v40
	v120 = v49 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v120
	v124 = v120
	v127 = v44
	v130 = v40
	goto L34
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v201
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v198
	return
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v198 = v178
	v201 = int32(0)
	goto L31
L33:
	;
	v156 = v127
	goto L38
L34:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v140 != 0 {
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v178 = v120 + v46
	goto L32
L36:
	;
	v141 = int32(1)
	v142 = v130 - v141
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v142
	if v141 < v127 {
		v124 = v124 + int32(2)
		v127 = v127 - v141
		v130 = v142
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124-int32(2)+v156<<(uint(int32(1))%32)))))
	if v172 != 0 {
		v198 = v124
		v201 = v156
		goto L31
	} else {
		goto L40
	}
L39:
	;
	v178 = v124
	goto L32
L40:
	;
	v173 = int32(1)
	if v173 < v156 {
		v156 = v156 - v173
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
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
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = v4
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 != int32(319) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v149 = v26 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v149 < v150 {
		v26 = v149
		goto L4
	} else {
		goto L41
	}
L7:
	;
	if v33 == int32(6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v138 = F_find_placeholder_info(m, l0, v32)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L39
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v39 = F_find_base_rel(m, l0, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L36
	}
L13:
	;
	return
L14:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+8)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v43 = int32(0)
	if l2 == v43 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v96 != 0 {
		goto L6
	} else {
		goto L29
	}
L16:
	;
	v96 = int32(1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v42 == int32(0) {
		v87 = v43
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v96 = v87
	goto L15
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v53 < v52 {
		v87 = v43
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(1)
	if v52 <= v55 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = v55
	goto L24
L23:
	;
	v58 = v52
	goto L24
L24:
	;
	v59 = int32(8)
	v64 = int32(0)
	goto L25
L25:
	;
	v71 = v64 << (uint(int32(2)) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2+v59+v71)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+(v42+v59))))
	v78 = v73 & (v75 ^ int32(-1))
	v80 = base.B2i32(v78 == int32(0))
	if v78 != 0 {
		v87 = v80
		goto L19
	} else {
		goto L27
	}
L26:
	;
	v87 = v80
	goto L19
L27:
	;
	v82 = v64 + int32(1)
	if v82 != v58 {
		v64 = v82
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+80)))
	v100 = (v41 - v97) << (uint(int32(2)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100+v101)))
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v118 = v103
	goto L32
L31:
	;
	v104 = F_copyObjectImpl(m, v32)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v119 = F_bms_add_members(m, v118, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v110 = F_lappend(m, v109, v104)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114+v100)))
	v118 = v116
	goto L32
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v100))) = v119
	goto L6
L36:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v128
	F_errmsg_internal(m, int32(477745), v12)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(488197), int32(329), int32(72186))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
	v141 = F_bms_add_members(m, v140, l2)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+20)) = v141
	goto L6
L41:
	;
	goto L5
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
						v41 = int32(0)
						v42 = F_range_cmp_bounds(m, l0, l2, l3)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = int32(0)
							v45 = base.B2i32(v44 <= v42)
							if v45&base.B2i32(v40 < v44) != 0 {
								v82 = v41
								m.G0 = v10 + int32(32)
								return v82
							} else {
								v49 = int32(0)
								if base.B2i32(v40 <= v49)|v45 == v49 {
									v82 = v41
									m.G0 = v10 + int32(32)
									return v82
								} else {
									v57 = F_range_cmp_bounds(m, l0, l1, l2)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
										if v59 == int32(1) {
											if v57 < int32(0) {
												v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
												*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v65
												v67 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
												*(*int64)(unsafe.Add(mBase, uint32(v10))) = v67
												v71 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													if v71 == int32(0) {
														v82 = int32(-1)
													} else {
														v82 = int32(1)
													}
													m.G0 = v10 + int32(32)
													return v82
												}
											} else {
												v82 = int32(1)
												m.G0 = v10 + int32(32)
												return v82
											}
										} else {
											if v57 <= int32(0) {
												v80 = int32(-1)
											} else {
												v80 = int32(1)
											}
											v82 = v80
											m.G0 = v10 + int32(32)
											return v82
										}
									}
								}
							}
						}
					}
				} else {
					v40 = int32(1)
					v41 = int32(0)
					v42 = F_range_cmp_bounds(m, l0, l2, l3)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(0)
						v45 = base.B2i32(v44 <= v42)
						if v45&base.B2i32(v40 < v44) != 0 {
							v82 = v41
							m.G0 = v10 + int32(32)
							return v82
						} else {
							v49 = int32(0)
							if base.B2i32(v40 <= v49)|v45 == v49 {
								v82 = v41
								m.G0 = v10 + int32(32)
								return v82
							} else {
								v57 = F_range_cmp_bounds(m, l0, l1, l2)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
									if v59 == int32(1) {
										if v57 < int32(0) {
											v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v65
											v67 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(v10))) = v67
											v71 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												if v71 == int32(0) {
													v82 = int32(-1)
												} else {
													v82 = int32(1)
												}
												m.G0 = v10 + int32(32)
												return v82
											}
										} else {
											v82 = int32(1)
											m.G0 = v10 + int32(32)
											return v82
										}
									} else {
										if v57 <= int32(0) {
											v80 = int32(-1)
										} else {
											v80 = int32(1)
										}
										v82 = v80
										m.G0 = v10 + int32(32)
										return v82
									}
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
				v41 = int32(0)
				v42 = F_range_cmp_bounds(m, l0, l2, l3)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = int32(0)
					v45 = base.B2i32(v44 <= v42)
					if v45&base.B2i32(v40 < v44) != 0 {
						v82 = v41
						m.G0 = v10 + int32(32)
						return v82
					} else {
						v49 = int32(0)
						if base.B2i32(v40 <= v49)|v45 == v49 {
							v82 = v41
							m.G0 = v10 + int32(32)
							return v82
						} else {
							v57 = F_range_cmp_bounds(m, l0, l1, l2)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
								if v59 == int32(1) {
									if v57 < int32(0) {
										v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v65
										v67 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
										*(*int64)(unsafe.Add(mBase, uint32(v10))) = v67
										v71 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											if v71 == int32(0) {
												v82 = int32(-1)
											} else {
												v82 = int32(1)
											}
											m.G0 = v10 + int32(32)
											return v82
										}
									} else {
										v82 = int32(1)
										m.G0 = v10 + int32(32)
										return v82
									}
								} else {
									if v57 <= int32(0) {
										v80 = int32(-1)
									} else {
										v80 = int32(1)
									}
									v82 = v80
									m.G0 = v10 + int32(32)
									return v82
								}
							}
						}
					}
				}
			}
		}
	} else {
		v57 = F_range_cmp_bounds(m, l0, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v59 == int32(1) {
				if v57 < int32(0) {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v65
					v67 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v67
					v71 = F_bounds_adjacent(m, l0, v10+int32(8), v10)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						if v71 == int32(0) {
							v82 = int32(-1)
						} else {
							v82 = int32(1)
						}
						m.G0 = v10 + int32(32)
						return v82
					}
				} else {
					v82 = int32(1)
					m.G0 = v10 + int32(32)
					return v82
				}
			} else {
				if v57 <= int32(0) {
					v80 = int32(-1)
				} else {
					v80 = int32(1)
				}
				v82 = v80
				m.G0 = v10 + int32(32)
				return v82
			}
		}
	}
}
func F_adjust_child_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = int32(0)
	if l1 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	goto L3
L3:
	;
	v13 = v4
	v14 = v4
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2+v14<<(uint(int32(2))%32))))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = F_bms_is_member(m, v20, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	return int32(0)
L7:
	;
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v34 = v13
	goto L10
L10:
	;
	v36 = v14 + int32(1)
	if v36 != l1 {
		v13 = v34
		v14 = v36
		goto L4
	} else {
		goto L17
	}
L11:
	;
	v27 = v13
	goto L13
L12:
	;
	v25 = F_bms_copy(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v29 = F_bms_del_member(m, v27, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	v27 = v25
	goto L13
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v32 = F_bms_add_member(m, v29, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v34 = v32
	goto L10
L17:
	;
	goto L5
L18:
	;
	v38 = v34
	goto L20
L19:
	;
	v38 = l0
	goto L20
L20:
	;
	return v38
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(23800)
			F_errmsg(m, int32(189571), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(485687), int32(175), int32(35650))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_anytime_typmod_check(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if int32(0) <= l1 {
		if base.Ui32(l1) < base.Ui32(int32(7)) {
			v42 = l1
			m.G0 = v7 + int32(32)
			return v42
		} else {
			v13 = int32(6)
			v16 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					v42 = v13
					m.G0 = v7 + int32(32)
					return v42
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(6)
						if l0 != 0 {
							v29 = int32(530611)
						} else {
							v29 = int32(731620)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v29
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
						F_errmsg(m, int32(479889), v7+int32(16))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490092), int32(85), int32(312970))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = v13
								m.G0 = v7 + int32(32)
								return v42
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if l0 != 0 {
					v56 = int32(530611)
				} else {
					v56 = int32(731620)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(337474), v7)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490092), int32(78), int32(312970))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
	var v16 int32
	_ = v16
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
	var v39 int32
	_ = v39
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
	v16 = v3
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v16<<(uint(int32(2))%32))))
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
	v74 = v16 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 < v75 {
		v14 = v65
		v16 = v74
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
	v39 = int32(0)
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34+v39<<(uint(int32(2))%32))))
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
	v51 = v39 + int32(1)
	if v31 != v51 {
		v39 = v51
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v9 = int32(0)
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
	if v15 <= v9 {
		v21 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v17 + v9<<(uint(int32(2))%32)
	goto L3
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v39
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v41)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+26)) = uint8(v43)
	v9 = v9 + int32(1)
	goto L1
L7:
	;
	return
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 <= v9 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v21 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = v28 + v9<<(uint(int32(2))%32)
	if v31 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
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
						F_errmsg(m, int32(17583), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errfinish(m, int32(485821), int32(1154), int32(159972))
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
										F_errmsg(m, int32(129052), v8+int32(16))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											F_errfinish(m, int32(485821), int32(1174), int32(159972))
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
								F_errmsg(m, int32(279285), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									F_errfinish(m, int32(485821), int32(1184), int32(159972))
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
				F_errmsg_internal(m, int32(42486), v8)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					F_errfinish(m, int32(485821), int32(1145), int32(159972))
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
	*(*int32)(unsafe.Add(mBase, _consts[258])) = v4
	return
}
func F_assign_io_method(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[901])))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v8
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
	*(*int32)(unsafe.Add(mBase, _consts[294])) = l0
	v6 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	if v6 == int32(13) {
		v9 = int32(4366432)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[291]))
		*(*int32)(unsafe.Add(mBase, _consts[291])) = v11 + int32(1)
	} else {
	}
	return
}
func F_assign_session_replication_role(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v4 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	if l0 != v4 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[459]))
		if v10 == int32(0) {
		} else {
			if v10 == int32(4089312) {
			} else {
				v15 = v10
				for {
					v19 = v15 - int32(5)
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
					if v20 != int32(1) {
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(96))))
						if v25 != 0 {
							v26 = F_stmt_requires_parse_analysis(m, v25)
							mBase = m.M
							if v26 != 0 {
								v36 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(12))))
								if v40 == v36 {
								} else {
									v43 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v40)+10)) = uint8(v43)
								}
							} else {
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(92))))
							if v29 == int32(0) {
							} else {
								v32 = F_query_requires_rewrite_plan(m, v29)
								mBase = m.M
								if v32 == int32(0) {
								} else {
									v36 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(12))))
									if v40 == v36 {
									} else {
										v43 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v40)+10)) = uint8(v43)
									}
								}
							}
						}
					}
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					if v47 != int32(4089312) {
						v15 = v47
						continue
					} else {
						break
					}
					break
				}
			}
		}
		v54 = *(*int32)(unsafe.Add(mBase, _consts[460]))
		if v54 == int32(0) {
		} else {
			if v54 == int32(4089320) {
			} else {
				v59 = v54
				for {
					v64 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v59-int32(16)))) = uint8(v64)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
					if v66 != int32(4089320) {
						v59 = v66
						continue
					} else {
						break
					}
					break
				}
			}
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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1414]))
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
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v8 == int32(0) {
		v27 = v7
		v28 = v8
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1415])))
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	if v28-v27 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v7 != v8 {
		v27 = v7
		v28 = v8
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v12 = v4
	v13 = l0
	goto L9
L9:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v27 = v16
		v28 = v17
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v27 = v16
	v28 = v17
	goto L6
L11:
	;
	v20 = int32(1)
	if v16 == v17 {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	goto L4
L14:
	;
	v35 = m.G0
	v36 = int32(16)
	v37 = v35 - v36
	m.G0 = v37
	v39 = int32(4364928)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[1416]))
	v41 = F_close(m, v40)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[1416])) = int32(-1)
	m.G0 = v37 + v36
	goto L17
L15:
	;
	v54 = v4
	goto L16
L16:
	;
	F_emscripten_builtin_free(m, v54)
	mBase = m.M
	v59 = F_strlen(m, l0)
	mBase = m.M
	v61 = v59 + int32(1)
	v62 = F_emscripten_builtin_malloc(m, v61)
	mBase = m.M
	if v62 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1415])) = uint8(v50)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[1414]))
	v54 = v53
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1414])) = v67
	goto L1
L19:
	;
	v67 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v66 = F___memcpy(m, v62, l0, v61)
	mBase = m.M
	v67 = v66
	goto L18
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
	var v106 float64
	_ = v106
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
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
					v106 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1560])))
					v111 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1561])))
					v114 = base.F64_sub(v106, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v111), v61))
					if v8 < int64(0) {
						v118 = base.F64_neg(v114)
					} else {
						v118 = v114
					}
					v119 = v118
					return v119
				}
			} else {
				v119 = l0
				return v119
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
				v106 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1560])))
				v111 = *(*float64)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[1561])))
				v114 = base.F64_sub(v106, base.F64_sub(base.F64_sub(base.F64_mul(v61, base.F64_add(v78, v95)), v111), v61))
				if v8 < int64(0) {
					v118 = base.F64_neg(v114)
				} else {
					v118 = v114
				}
				v119 = v118
				return v119
			}
		}
	}
}
func F_atan2(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v106 float64
	_ = v106
	var v123 float64
	_ = v123
	var v130 int32
	_ = v130
	var v133 float64
	_ = v133
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v158 float64
	_ = v158
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v200 float64
	_ = v200
	var v218 float64
	_ = v218
	var v225 int64
	_ = v225
	var v230 int32
	_ = v230
	var v240 float64
	_ = v240
	var v246 float64
	_ = v246
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v294 float64
	_ = v294
	var v311 float64
	_ = v311
	var v318 int32
	_ = v318
	var v321 float64
	_ = v321
	var v326 float64
	_ = v326
	var v329 float64
	_ = v329
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v346 float64
	_ = v346
	var v347 float64
	_ = v347
	var v366 float64
	_ = v366
	var v367 float64
	_ = v367
	if base.Ui64(base.I64_reinterpret_f64(l1)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v21 = base.I64_reinterpret_f64(l1)
			v24 = base.I32_wrap_i64(int64(base.Ui64(v21) >> (uint(int64(32)) % 64)))
			v27 = base.I32_wrap_i64(v21)
			if v24-int32(1072693248)|v27 == int32(0) {
				v37 = base.I64_reinterpret_f64(l0)
				v42 = base.I32_wrap_i64(int64(base.Ui64(v37)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1141899264)) <= base.Ui32(v42) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l0)&int64(9223372036854775807)) {
						v52 = l0
					} else {
						v52 = base.F64_copysign(float64(1.5707963267948966), l0)
					}
					v158 = v52
				} else {
					if base.Ui32(v42) <= base.Ui32(int32(1071382527)) {
						if base.Ui32(int32(1044381696)) <= base.Ui32(v42) {
							v89 = l0
							v90 = int32(-1)
							v91 = base.F64_mul(v89, v89)
							v92 = base.F64_mul(v91, v91)
							v106 = base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
							v123 = base.F64_mul(v91, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
							if base.Ui32(v42) <= base.Ui32(int32(1071382527)) {
								v158 = base.F64_sub(v89, base.F64_mul(v89, base.F64_add(v106, v123)))
							} else {
								v130 = v90 << (uint(int32(3)) % 32)
								v133 = *(*float64)(unsafe.Add(mBase, uint32(v130)+uint32(_consts[1560])))
								v138 = *(*float64)(unsafe.Add(mBase, uint32(v130)+uint32(_consts[1561])))
								v141 = base.F64_sub(v133, base.F64_sub(base.F64_sub(base.F64_mul(v89, base.F64_add(v106, v123)), v138), v89))
								if v37 < int64(0) {
									v145 = base.F64_neg(v141)
								} else {
									v145 = v141
								}
								v146 = v145
								v158 = v146
							}
						} else {
							v146 = l0
							v158 = v146
						}
					} else {
						v58 = base.F64_abs(l0)
						if base.Ui32(v42) <= base.Ui32(int32(1072889855)) {
							if base.Ui32(v42) <= base.Ui32(int32(1072037887)) {
								v89 = base.F64_div(base.F64_add(base.F64_add(v58, v58), float64(-1)), base.F64_add(v58, float64(2)))
								v90 = int32(0)
							} else {
								v89 = base.F64_div(base.F64_add(v58, float64(-1)), base.F64_add(v58, float64(1)))
								v90 = int32(1)
							}
						} else {
							if base.Ui32(v42) <= base.Ui32(int32(1073971199)) {
								v89 = base.F64_div(base.F64_add(v58, float64(-1.5)), base.F64_add(base.F64_mul(v58, float64(1.5)), float64(1)))
								v90 = int32(2)
							} else {
								v89 = base.F64_div(float64(-1), v58)
								v90 = int32(3)
							}
						}
						v91 = base.F64_mul(v89, v89)
						v92 = base.F64_mul(v91, v91)
						v106 = base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
						v123 = base.F64_mul(v91, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, base.F64_add(base.F64_mul(v92, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
						if base.Ui32(v42) <= base.Ui32(int32(1071382527)) {
							v158 = base.F64_sub(v89, base.F64_mul(v89, base.F64_add(v106, v123)))
						} else {
							v130 = v90 << (uint(int32(3)) % 32)
							v133 = *(*float64)(unsafe.Add(mBase, uint32(v130)+uint32(_consts[1560])))
							v138 = *(*float64)(unsafe.Add(mBase, uint32(v130)+uint32(_consts[1561])))
							v141 = base.F64_sub(v133, base.F64_sub(base.F64_sub(base.F64_mul(v89, base.F64_add(v106, v123)), v138), v89))
							if v37 < int64(0) {
								v145 = base.F64_neg(v141)
							} else {
								v145 = v141
							}
							v146 = v145
							v158 = v146
						}
					}
				}
				return v158
			} else {
				v163 = int32(base.Ui32(v24)>>(uint(int32(30))%32)) & int32(2)
				v164 = base.I64_reinterpret_f64(l0)
				v168 = v163 | base.I32_wrap_i64(int64(base.Ui64(v164)>>(uint(int64(63))%64)))
				v173 = base.I32_wrap_i64(int64(base.Ui64(v164)>>(uint(int64(32))%64))) & int32(2147483647)
				if v173|base.I32_wrap_i64(v164) == int32(0) {
					switch v168 - int32(2) {
					case 0:
						return float64(3.141592653589793)
					case 1:
						return float64(-3.141592653589793)
					default:
						v367 = l0
						return v367
					}
				} else {
					v185 = v24 & int32(2147483647)
					if v185|v27 == int32(0) {
						return base.F64_copysign(float64(1.5707963267948966), l0)
					} else {
						if v185 == int32(2146435072) {
							if v173 != int32(2146435072) {
								v366 = *(*float64)(unsafe.Add(mBase, uint32(v168<<(uint(int32(3))%32))+uint32(_consts[1562])))
								v367 = v366
								return v367
							} else {
								v200 = *(*float64)(unsafe.Add(mBase, uint32(v168<<(uint(int32(3))%32))+uint32(_consts[1563])))
								return v200
							}
						} else {
							if base.B2i32(v173 != int32(2146435072))&base.B2i32(base.Ui32(v173) <= base.Ui32(v185+int32(67108864))) == int32(0) {
								return base.F64_copysign(float64(1.5707963267948966), l0)
							} else {
								if v163 != 0 {
									if base.Ui32(v173+int32(67108864)) < base.Ui32(v185) {
										v347 = float64(0)
									} else {
										v218 = base.F64_abs(base.F64_div(l0, l1))
										v225 = base.I64_reinterpret_f64(v218)
										v230 = base.I32_wrap_i64(int64(base.Ui64(v225)>>(uint(int64(32))%64))) & int32(2147483647)
										if base.Ui32(int32(1141899264)) <= base.Ui32(v230) {
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v218)&int64(9223372036854775807)) {
												v240 = v218
											} else {
												v240 = base.F64_copysign(float64(1.5707963267948966), v218)
											}
											v346 = v240
										} else {
											if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
												if base.Ui32(int32(1044381696)) <= base.Ui32(v230) {
													v277 = v218
													v278 = int32(-1)
													v279 = base.F64_mul(v277, v277)
													v280 = base.F64_mul(v279, v279)
													v294 = base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
													v311 = base.F64_mul(v279, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
													if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
														v346 = base.F64_sub(v277, base.F64_mul(v277, base.F64_add(v294, v311)))
													} else {
														v318 = v278 << (uint(int32(3)) % 32)
														v321 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1560])))
														v326 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1561])))
														v329 = base.F64_sub(v321, base.F64_sub(base.F64_sub(base.F64_mul(v277, base.F64_add(v294, v311)), v326), v277))
														if v225 < int64(0) {
															v333 = base.F64_neg(v329)
														} else {
															v333 = v329
														}
														v334 = v333
														v346 = v334
													}
												} else {
													v334 = v218
													v346 = v334
												}
											} else {
												v246 = base.F64_abs(v218)
												if base.Ui32(v230) <= base.Ui32(int32(1072889855)) {
													if base.Ui32(v230) <= base.Ui32(int32(1072037887)) {
														v277 = base.F64_div(base.F64_add(base.F64_add(v246, v246), float64(-1)), base.F64_add(v246, float64(2)))
														v278 = int32(0)
													} else {
														v277 = base.F64_div(base.F64_add(v246, float64(-1)), base.F64_add(v246, float64(1)))
														v278 = int32(1)
													}
												} else {
													if base.Ui32(v230) <= base.Ui32(int32(1073971199)) {
														v277 = base.F64_div(base.F64_add(v246, float64(-1.5)), base.F64_add(base.F64_mul(v246, float64(1.5)), float64(1)))
														v278 = int32(2)
													} else {
														v277 = base.F64_div(float64(-1), v246)
														v278 = int32(3)
													}
												}
												v279 = base.F64_mul(v277, v277)
												v280 = base.F64_mul(v279, v279)
												v294 = base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
												v311 = base.F64_mul(v279, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
												if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
													v346 = base.F64_sub(v277, base.F64_mul(v277, base.F64_add(v294, v311)))
												} else {
													v318 = v278 << (uint(int32(3)) % 32)
													v321 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1560])))
													v326 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1561])))
													v329 = base.F64_sub(v321, base.F64_sub(base.F64_sub(base.F64_mul(v277, base.F64_add(v294, v311)), v326), v277))
													if v225 < int64(0) {
														v333 = base.F64_neg(v329)
													} else {
														v333 = v329
													}
													v334 = v333
													v346 = v334
												}
											}
										}
										v347 = v346
									}
								} else {
									v218 = base.F64_abs(base.F64_div(l0, l1))
									v225 = base.I64_reinterpret_f64(v218)
									v230 = base.I32_wrap_i64(int64(base.Ui64(v225)>>(uint(int64(32))%64))) & int32(2147483647)
									if base.Ui32(int32(1141899264)) <= base.Ui32(v230) {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v218)&int64(9223372036854775807)) {
											v240 = v218
										} else {
											v240 = base.F64_copysign(float64(1.5707963267948966), v218)
										}
										v346 = v240
									} else {
										if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
											if base.Ui32(int32(1044381696)) <= base.Ui32(v230) {
												v277 = v218
												v278 = int32(-1)
												v279 = base.F64_mul(v277, v277)
												v280 = base.F64_mul(v279, v279)
												v294 = base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
												v311 = base.F64_mul(v279, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
												if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
													v346 = base.F64_sub(v277, base.F64_mul(v277, base.F64_add(v294, v311)))
												} else {
													v318 = v278 << (uint(int32(3)) % 32)
													v321 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1560])))
													v326 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1561])))
													v329 = base.F64_sub(v321, base.F64_sub(base.F64_sub(base.F64_mul(v277, base.F64_add(v294, v311)), v326), v277))
													if v225 < int64(0) {
														v333 = base.F64_neg(v329)
													} else {
														v333 = v329
													}
													v334 = v333
													v346 = v334
												}
											} else {
												v334 = v218
												v346 = v334
											}
										} else {
											v246 = base.F64_abs(v218)
											if base.Ui32(v230) <= base.Ui32(int32(1072889855)) {
												if base.Ui32(v230) <= base.Ui32(int32(1072037887)) {
													v277 = base.F64_div(base.F64_add(base.F64_add(v246, v246), float64(-1)), base.F64_add(v246, float64(2)))
													v278 = int32(0)
												} else {
													v277 = base.F64_div(base.F64_add(v246, float64(-1)), base.F64_add(v246, float64(1)))
													v278 = int32(1)
												}
											} else {
												if base.Ui32(v230) <= base.Ui32(int32(1073971199)) {
													v277 = base.F64_div(base.F64_add(v246, float64(-1.5)), base.F64_add(base.F64_mul(v246, float64(1.5)), float64(1)))
													v278 = int32(2)
												} else {
													v277 = base.F64_div(float64(-1), v246)
													v278 = int32(3)
												}
											}
											v279 = base.F64_mul(v277, v277)
											v280 = base.F64_mul(v279, v279)
											v294 = base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
											v311 = base.F64_mul(v279, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, base.F64_add(base.F64_mul(v280, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
											if base.Ui32(v230) <= base.Ui32(int32(1071382527)) {
												v346 = base.F64_sub(v277, base.F64_mul(v277, base.F64_add(v294, v311)))
											} else {
												v318 = v278 << (uint(int32(3)) % 32)
												v321 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1560])))
												v326 = *(*float64)(unsafe.Add(mBase, uint32(v318)+uint32(_consts[1561])))
												v329 = base.F64_sub(v321, base.F64_sub(base.F64_sub(base.F64_mul(v277, base.F64_add(v294, v311)), v326), v277))
												if v225 < int64(0) {
													v333 = base.F64_neg(v329)
												} else {
													v333 = v329
												}
												v334 = v333
												v346 = v334
											}
										}
									}
									v347 = v346
								}
								switch v168 - int32(1) {
								case 0:
									return base.F64_neg(v347)
								case 1:
									return base.F64_sub(float64(3.141592653589793), base.F64_add(v347, float64(-1.2246467991473532e-16)))
								case 2:
									return base.F64_add(base.F64_add(v347, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
								default:
									v367 = v347
									return v367
								}
							}
						}
					}
				}
			}
		} else {
			return base.F64_add(l0, l1)
		}
	} else {
		return base.F64_add(l0, l1)
	}
}
