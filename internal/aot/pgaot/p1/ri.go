package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_check_ins(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_check_ins_0), int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_RI_FKey_check(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_noaction_del(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_noaction_del_0), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_setdefault_del(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setdefault_del_0), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(0), int32(3))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_ri_FetchConstraintInfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v11 != 0 {
		v12 = F_ri_LoadConstraintInfo(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
			if l2 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 == v16 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
					if v19 == v20 {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)))
						switch v47 - int32(102) {
						case 0, 13:
							m.G0 = v9 - int32(-64)
							return v12
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+164)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
								F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_0), v7+int32(-48))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2345), int32(_a_F_ri_FetchConstraintInfo_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 10:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_ri_FetchConstraintInfo_3), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2350), int32(_a_F_ri_FetchConstraintInfo_2))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v27
							*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v26 + int32(4)
							F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_4), v7+int32(-32))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2331), int32(_a_F_ri_FetchConstraintInfo_2))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v27
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v26 + int32(4)
						F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_4), v7+int32(-32))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2331), int32(_a_F_ri_FetchConstraintInfo_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
				if v16 != v42 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v116
						*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v115 + int32(4)
						F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_4), v7+int32(-16))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2338), int32(_a_F_ri_FetchConstraintInfo_2))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v44 != v45 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v116
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v115 + int32(4)
							F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_4), v7+int32(-16))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2338), int32(_a_F_ri_FetchConstraintInfo_2))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)))
						switch v47 - int32(102) {
						case 0, 13:
							m.G0 = v9 - int32(-64)
							return v12
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+164)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v54
								F_errmsg_internal(m, int32(_a_F_ri_FetchConstraintInfo_0), v7+int32(-48))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2345), int32(_a_F_ri_FetchConstraintInfo_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 10:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_ri_FetchConstraintInfo_3), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2350), int32(_a_F_ri_FetchConstraintInfo_2))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(117833860))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v93 + int32(4)
				F_errmsg(m, int32(_a_F_ri_FetchConstraintInfo_5), v9)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_ri_FetchConstraintInfo_6), int32(0))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ri_FetchConstraintInfo_1), int32(2320), int32(_a_F_ri_FetchConstraintInfo_2))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
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
}
