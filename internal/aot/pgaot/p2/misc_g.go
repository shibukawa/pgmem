package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAccessStrategyWithSize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	if base.Ui32(int32(15)) <= base.Ui32(l0+int32(7)) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[133]))
		v9 = int32(8)
		v10 = base.I32_div_s(v8, v9)
		v12 = base.I32_div_s(l0, v9)
		if v10 < v12 {
			v14 = v10
		} else {
			v14 = v12
		}
		v19 = F_palloc0(m, v14<<(uint(int32(2))%32)+int32(12))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(3)
			v27 = v19
			return v27
		}
	} else {
		v27 = int32(0)
		return v27
	}
}
func F_GetComboCommandId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1244]))
	if v12 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[68]))
		v18 = F_MemoryContextAlloc(m, v16, int32(800))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = int32(100)
			*(*int32)(unsafe.Add(mBase, _consts[1245])) = v23
			*(*int32)(unsafe.Add(mBase, _consts[66])) = v18
			*(*int32)(unsafe.Add(mBase, _consts[1246])) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+28)) = int64(51539607560)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[68]))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v33
			v41 = F_hash_create(m, int32(164962), v23, v7+int32(-52), int32(1064))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1244])) = v41
				v44 = v41
				v46 = *(*int32)(unsafe.Add(mBase, _consts[1245]))
				v48 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
				if v46 <= v48 {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[66]))
					v54 = F_repalloc(m, v51, v46<<(uint(int32(4))%32))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1245])) = v46 << (uint(int32(1)) % 32)
						*(*int32)(unsafe.Add(mBase, _consts[66])) = v54
						v63 = *(*int32)(unsafe.Add(mBase, _consts[1244]))
						v64 = v63
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
						v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
							if v74 == int32(1) {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
								v92 = v77
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _consts[66]))
								v80 = int32(4442660)
								v81 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
								v84 = v79 + v81<<(uint(int32(3))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
								*(*int32)(unsafe.Add(mBase, _consts[1246])) = v81 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
								v92 = v81
							}
							m.G0 = v9 - int32(-64)
							return v92
						}
					}
				} else {
					v64 = v44
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
					v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
						if v74 == int32(1) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
							v92 = v77
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, _consts[66]))
							v80 = int32(4442660)
							v81 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
							v84 = v79 + v81<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
							*(*int32)(unsafe.Add(mBase, _consts[1246])) = v81 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
							v92 = v81
						}
						m.G0 = v9 - int32(-64)
						return v92
					}
				}
			}
		}
	} else {
		v44 = v12
		v46 = *(*int32)(unsafe.Add(mBase, _consts[1245]))
		v48 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
		if v46 <= v48 {
			v51 = *(*int32)(unsafe.Add(mBase, _consts[66]))
			v54 = F_repalloc(m, v51, v46<<(uint(int32(4))%32))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1245])) = v46 << (uint(int32(1)) % 32)
				*(*int32)(unsafe.Add(mBase, _consts[66])) = v54
				v63 = *(*int32)(unsafe.Add(mBase, _consts[1244]))
				v64 = v63
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
				v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
					if v74 == int32(1) {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
						v92 = v77
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _consts[66]))
						v80 = int32(4442660)
						v81 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
						v84 = v79 + v81<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
						*(*int32)(unsafe.Add(mBase, _consts[1246])) = v81 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
						v92 = v81
					}
					m.G0 = v9 - int32(-64)
					return v92
				}
			}
		} else {
			v64 = v44
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l0
			v72 = F_hash_search(m, v64, v7+int32(-52), int32(1), v7+int32(-1))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+63)))
				if v74 == int32(1) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
					v92 = v77
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _consts[66]))
					v80 = int32(4442660)
					v81 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
					v84 = v79 + v81<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
					*(*int32)(unsafe.Add(mBase, _consts[1246])) = v81 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v81
					v92 = v81
				}
				m.G0 = v9 - int32(-64)
				return v92
			}
		}
	}
}
func F_GetCompressionMethodName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	switch l0 - int32(108) {
	case 0:
		v27 = int32(526133)
		m.G0 = v6 + int32(16)
		return v27
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg_internal(m, int32(478105), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(473164), int32(313), int32(363330))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v27 = int32(7729)
		m.G0 = v6 + int32(16)
		return v27
	}
}
func F_GetFdwRoutineByServerId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(32), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15)+72))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = F_SearchSysCache1(m, int32(30), v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
							F_errmsg_internal(m, int32(41303), v7+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473392), int32(397), int32(443542))
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
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
						v27 = v25 + v26
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v27 + int32(4)
									F_errmsg(m, int32(208899), v7+int32(32))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(473392), int32(406), int32(443542))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
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
							F_ReleaseCatCache(m, v21)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = F_GetFdwRoutine(m, v28)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v33
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(41143), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(473392), int32(389), int32(443542))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_GetLatestXTime(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+96)) = int32(1)
	if v5 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[248]))
		F_s_lock(m, v9+int32(96), int32(469368), int32(4642), int32(358018))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[248]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(0)
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+64))
			return v23
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[248]))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(0)
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+64))
		return v23
	}
}
func F_GetOldestNonRemovableTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_ComputeXidHorizons(m, v5+int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		if l0 == v13 {
			v44 = v13
			v48 = v44
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
			if v18 != 0 {
				v44 = v13
				v48 = v44
			} else {
				v19 = F_RecoveryInProgress(m)
				mBase = m.M
				if v19 != 0 {
					v44 = v13
					v48 = v44
				} else {
					v20 = int32(1)
					v21 = F_IsCatalogRelation(m, l0)
					mBase = m.M
					if v21 != 0 {
						v44 = v20
						v48 = v44
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						if v23 < int32(2) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v40 != 0 {
								v44 = int32(3)
								v48 = v44
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								if v41 != 0 {
									v44 = int32(3)
									v48 = v44
								} else {
									v48 = int32(2)
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+118)))
							if v27 != int32(112) {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v40 != 0 {
									v44 = int32(3)
									v48 = v44
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v41 != 0 {
										v44 = int32(3)
										v48 = v44
									} else {
										v48 = int32(2)
									}
								}
							} else {
								v30 = F_IsCatalogRelation(m, l0)
								mBase = m.M
								if v30 != 0 {
									v44 = v20
									v48 = v44
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
									if v31 == int32(0) {
										v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
										if v40 != 0 {
											v44 = int32(3)
											v48 = v44
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v41 != 0 {
												v44 = int32(3)
												v48 = v44
											} else {
												v48 = int32(2)
											}
										}
									} else {
										v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+119)))
										switch v35 - int32(109) {
										case 0, 5:
											v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+104)))
											if v38 != 0 {
												v44 = v20
												v48 = v44
											} else {
												v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												if v40 != 0 {
													v44 = int32(3)
													v48 = v44
												} else {
													v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v41 != 0 {
														v44 = int32(3)
														v48 = v44
													} else {
														v48 = int32(2)
													}
												}
											}
										default:
											v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
											if v40 != 0 {
												v44 = int32(3)
												v48 = v44
											} else {
												v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v41 != 0 {
													v44 = int32(3)
													v48 = v44
												} else {
													v48 = int32(2)
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
		switch v48 - int32(1) {
		case 0:
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
			v55 = v52
		case 1:
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v5)+40))
			v55 = v53
		case 2:
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			v55 = v54
		default:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
			v55 = v51
		}
		m.G0 = v5 + int32(48)
		return v55
	}
}
func F_GetOldestSafeDecodingTransactionId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	v8 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v11 == int32(1) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+316))
		v19 = base.B2i32(v17 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v19)
		v21 = v19
	} else {
		v21 = int32(0)
	}
	v23 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v27 = F_LWLockAcquire(m, v23+int32(384), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, _consts[139]))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
		v35 = *(*int32)(unsafe.Add(mBase, _consts[193]))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
		if v36 == int32(0) {
			v56 = v33
		} else {
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v33))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v36)) == int32(0) {
				v50 = base.B2i32(base.Ui32(v36) < base.Ui32(v33))
			} else {
				v50 = int32(base.Ui32(v36-v33) >> (uint(int32(31)) % 32))
			}
			if v50 == int32(0) {
				v56 = v33
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, _consts[193]))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
				v56 = v55
			}
		}
		if l0 == int32(0) {
			v82 = v56
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, _consts[193]))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
			if v61 == int32(0) {
				v82 = v56
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v56))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v61)) == int32(0) {
					v75 = base.B2i32(base.Ui32(v61) < base.Ui32(v56))
				} else {
					v75 = int32(base.Ui32(v61-v56) >> (uint(int32(31)) % 32))
				}
				if v75 == int32(0) {
					v82 = v56
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _consts[193]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
					v82 = v80
				}
			}
		}
		if v21 != 0 {
			v122 = v82
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if v83 <= int32(0) {
				v122 = v82
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, _consts[153]))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
				v90 = int32(0)
				v91 = v82
				v93 = v83
				for {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90<<(uint(int32(2))%32))))
					if base.Ui32(int32(3)) <= base.Ui32(v99) {
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v99)) == int32(0) {
							v113 = base.B2i32(base.Ui32(v99) < base.Ui32(v91))
						} else {
							v113 = int32(base.Ui32(v99-v91) >> (uint(int32(31)) % 32))
						}
						if v113 != 0 {
							v114 = v99
						} else {
							v114 = v91
						}
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v116 = v114
						v117 = v115
					} else {
						v116 = v91
						v117 = v93
					}
					v119 = v90 + int32(1)
					if v119 < v117 {
						v90 = v119
						v91 = v116
						v93 = v117
						continue
					} else {
						break
					}
					break
				}
				v122 = v116
			}
		}
		v128 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		F_LWLockRelease(m, v128+int32(384))
		mBase = m.M
		v132 = m.ExcPending
		if v132 != 0 {
			return int32(0)
		} else {
			return v122
		}
	}
}
func F_GetOldestUnsummarizedLSN(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
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
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v203 int32
	_ = v203
	var v214 int64
	_ = v214
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v282 int64
	_ = v282
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[585])))
	if v23 != int32(1) {
		v282 = int64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v282
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	if v27 != int32(15) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v35 = F_LWLockAcquire(m, v31+int32(6272), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v64 = F_GetLatestLSN(m, v20+int32(12))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L19
	}
L6:
	;
	return int64(0)
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v41 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v56+int32(6272))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L18
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
	goto L13
L12:
	;
	goto L13
L13:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v47)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v50+int32(6272))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v282 = v44
	goto L1
L18:
	;
	goto L5
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v67 = F_readTimeLineHistory(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = v69
	goto L23
L22:
	;
	v71 = int32(0)
	goto L23
L23:
	;
	v74 = v71
	goto L25
L24:
	;
	v112 = int64(0)
	v114 = F_GetWalSummaries(m, v109, v112, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L33
	}
L25:
	;
	v90 = v74 - int32(1)
	if v90 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v106 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v109 = v104
	v111 = v100 * v106
	goto L24
L27:
	;
	v109 = v3
	v111 = int64(0)
	goto L24
L28:
	;
	goto L29
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v90<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = F_XLogGetOldestSegno(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v100 == int64(0) {
		v74 = v90
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	if v109 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L33:
	;
	if v114 == int32(0) {
		v203 = v3
		v214 = v111
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v118 <= int32(0) {
		v203 = v3
		v214 = v111
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v121 = int32(0)
	if v121 < v118 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v124 = v118
	goto L38
L37:
	;
	v124 = v121
	goto L38
L38:
	;
	v125 = int32(1)
	if v118 == v125 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v124&v125 == int32(0) {
		v203 = v175
		v214 = v186
		goto L32
	} else {
		goto L52
	}
L40:
	;
	v129 = int32(0)
	v174 = v129
	v175 = v129
	v186 = v111
	goto L39
L41:
	;
	goto L42
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v134 = int32(0)
	v139 = v134
	v140 = v134
	v142 = v134
	v151 = v111
	goto L43
L43:
	;
	v156 = v133 + v139<<(uint(int32(2))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	v161 = base.B2i32(base.Ui64(v151) < base.Ui64(v160))
	if base.Ui64(v151) < base.Ui64(v160) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v174 = v168
	v175 = v166
	v186 = v164
	goto L39
L45:
	;
	v162 = v160
	goto L47
L46:
	;
	v162 = v151
	goto L47
L47:
	;
	v163 = base.B2i32(base.Ui64(v162) < base.Ui64(v158))
	if base.Ui64(v162) < base.Ui64(v158) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v164 = v158
	goto L50
L49:
	;
	v164 = v162
	goto L50
L50:
	;
	v166 = v161 | v163 | v140
	v167 = int32(2)
	v168 = v139 + v167
	v170 = v142 + v167
	if v170 != v124&int32(2147483646) {
		v139 = v168
		v140 = v166
		v142 = v170
		v151 = v164
		goto L43
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v174<<(uint(int32(2))%32))))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)+8))
	v197 = base.B2i32(base.Ui64(v186) < base.Ui64(v196))
	if base.Ui64(v186) < base.Ui64(v196) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v198 = v196
	goto L55
L54:
	;
	v198 = v186
	goto L55
L55:
	;
	v203 = v197 | v175
	v214 = v198
	goto L32
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v241 = F_LWLockAcquire(m, v237+int32(6272), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L63
	}
L59:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v226
	F_errmsg_internal(m, int32(48323), v20)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(472018), int32(599), int32(504150))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
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
	v244 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	if v27 != int32(15) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if l0 != 0 {
		goto L70
	} else {
		goto L71
	}
L65:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v244)+8))
	v257 = v256
	goto L64
L66:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v247 != 0 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v244)+8)) = v214
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v244))) = uint8(v249)
	*(*int64)(unsafe.Add(mBase, uint32(v244)+24)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v109
	v254 = v203 & v249
	*(*uint8)(unsafe.Add(mBase, uint32(v244)+16)) = uint8(v254)
	v257 = v214
	goto L64
L69:
	;
	goto L68
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v258
	goto L72
L71:
	;
	goto L72
L72:
	;
	if l1 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v260)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v263+int32(6272))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	v282 = v257
	goto L1
}
func F_GetSerializableTransactionSnapshotInt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v281 int64
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v400 int32
	_ = v400
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v468 int64
	_ = v468
	var v470 int32
	_ = v470
	var v471 int64
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v26 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L13
	} else {
		goto L169
	}
L2:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v80)+64))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v80)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v810)+4)) = v811
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v80)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v811))) = v813
	v816 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	if v817 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L3:
	;
	if v29&int32(1) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v29 = int32(1)
	goto L6
L5:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+76)))
	v29 = v28
	goto L6
L6:
	;
	goto L3
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L13
	} else {
		goto L157
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v60 = F_LWLockAcquire(m, v56+int32(3584), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if l1 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v66 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v96+int32(3584))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L13
	} else {
		goto L22
	}
L16:
	;
	if v66 == v65 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v72 = v66 + int32(4)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v75
	v78 = v65 + int32(8)
	v80 = v66 + int32(-64)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if v81 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v78
	goto L20
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v78
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v66
	if v80 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v106 = F_LWLockAcquire(m, v102+int32(3712), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[805]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v110 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v419+int32(3712))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L13
	} else {
		goto L87
	}
L25:
	;
	if v110 == v109 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = int64(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v110)+40))
	if v123 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_ReleaseOneSerializableXact(m, v110-int32(56), int32(0), int32(1))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L13
	} else {
		goto L86
	}
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v110)+52))
	if v126&int32(32) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v126&int32(16) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v110-int32(32))))
	v135 = v134
	goto L32
L31:
	;
	v135 = int64(-1)
	goto L32
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+28))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, _consts[807])))
	v142 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v146 = F_LWLockAcquire(m, v142+int32(6656), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v149 = int32(base.Ui32(v123) >> (uint(int32(10)) % 32))
	v150 = base.I32_rem_u_s(v149, v140)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	if v153 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v375+int32(6656))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L85
	}
L35:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v153))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v123)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v167 != 0 {
		goto L34
	} else {
		goto L40
	}
L37:
	;
	v167 = base.B2i32(base.Ui32(v123) < base.Ui32(v153))
	goto L36
L38:
	;
	goto L39
L39:
	;
	v167 = int32(base.Ui32(v123-v153) >> (uint(int32(31)) % 32))
	goto L36
L40:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
	if v170 < int64(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	if v227 != 0 {
		goto L58
	} else {
		goto L59
	}
L42:
	;
	v221 = int32(1)
	v224 = base.I64_extend_i32_u(int32(base.Ui32(v153) >> (uint(int32(10)) % 32)))
	goto L41
L43:
	;
	goto L44
L44:
	;
	if base.Ui64(v170) <= base.Ui64(int64(4194302)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v182 = v170 + int64(1)
	goto L47
L46:
	;
	v182 = int64(0)
	goto L47
L47:
	;
	v183 = int32(0)
	v187 = int32(4)
	v188 = base.I32_wrap_i64(v170)<<(uint(int32(10))%32) | v187
	v190 = v123 & int32(-1024)
	v192 = v190 | v187
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v192))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == v183 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v204 == int32(0) {
		v221 = v183
		v224 = v182
		goto L41
	} else {
		goto L52
	}
L49:
	;
	v204 = base.B2i32(base.Ui32(v188) < base.Ui32(v192))
	goto L48
L50:
	;
	goto L51
L51:
	;
	v204 = int32(base.Ui32(v188-v192) >> (uint(int32(31)) % 32))
	goto L48
L52:
	;
	v208 = v190 + int32(1027)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v208))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v188)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v221 = v220
	v224 = v182
	goto L41
L54:
	;
	v220 = base.B2i32(base.Ui32(v188) < base.Ui32(v208))
	goto L53
L55:
	;
	goto L56
L56:
	;
	v220 = int32(base.Ui32(v188-v208) >> (uint(int32(31)) % 32))
	goto L53
L57:
	;
	v246 = base.I64_extend_i32_u(v149)
	if v221 != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v227))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v123)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v244 = v226
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v123
	goto L57
L61:
	;
	if v239 == int32(0) {
		goto L57
	} else {
		goto L65
	}
L62:
	;
	v239 = base.B2i32(base.Ui32(v227) < base.Ui32(v123))
	goto L61
L63:
	;
	goto L64
L64:
	;
	v239 = base.B2i32(int32(0) < v123-v227)
	goto L61
L65:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v244 = v243
	goto L60
L66:
	;
	v336 = int32(4365844)
	v337 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v325<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v342+v123<<(uint(int32(3))%32)&int32(8184)))) = v135
	v350 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v353 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v351+v325))) = uint8(v353)
	F_LWLockRelease(m, v322)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L13
	} else {
		goto L84
	}
L67:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	*(*int64)(unsafe.Add(mBase, uint32(v248))) = v246
	v251 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+28))
	v254 = int64(*(*uint16)(unsafe.Add(mBase, _consts[807])))
	v255 = base.I64_rem_u_s(v224, v254)
	v259 = v252 + base.I32_wrap_i64(v255)<<(uint(int32(7))%32)
	v261 = F_LWLockAcquire(m, v259, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L13
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v311 = v138 + v150<<(uint(int32(7))%32)
	v313 = F_LWLockAcquire(m, v311, int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L13
	} else {
		goto L82
	}
L70:
	;
	v264 = F_SimpleLruZeroPage(m, int32(4365844), v224)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	if v224 == v246 {
		v322 = v259
		v325 = v264
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v270 = v259
	v281 = v224
	goto L73
L73:
	;
	F_LWLockRelease(m, v270)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L75
	}
L74:
	;
	v322 = v301
	v325 = v306
	goto L66
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+28))
	if v281 <= int64(4194302) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v294 = v281 + int64(1)
	goto L78
L77:
	;
	v294 = int64(0)
	goto L78
L78:
	;
	v296 = int64(*(*uint16)(unsafe.Add(mBase, _consts[807])))
	v297 = base.I64_rem_s(v294, v296)
	v301 = v288 + base.I32_wrap_i64(v297)<<(uint(int32(7))%32)
	v303 = F_LWLockAcquire(m, v301, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	v306 = F_SimpleLruZeroPage(m, int32(4365844), v294)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	if v294 != v246 {
		v270 = v301
		v281 = v294
		goto L73
	} else {
		goto L81
	}
L81:
	;
	goto L74
L82:
	;
	v317 = F_SimpleLruReadPage(m, int32(4365844), v246, int32(1), v123)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	v322 = v311
	v325 = v317
	goto L66
L84:
	;
	goto L34
L85:
	;
	goto L27
L86:
	;
	goto L24
L87:
	;
	goto L10
L88:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, _consts[809])))
	if v437 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v426 = F_GetSnapshotData(m, l0)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v429 = F_ProcArrayInstallImportedXmin(m, v428, l1)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L13
	} else {
		goto L93
	}
L92:
	;
	v433 = v426
	goto L88
L93:
	;
	if v429 == int32(0) {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v433 = l0
	goto L88
L95:
	;
	m.G0 = v20 - int32(-64)
	return v433
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v66-int32(60)))) = v36
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v435)+32))
	v470 = v66 - int32(56)
	v471 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v470))) = v471
	*(*int64)(unsafe.Add(mBase, uint32(v66-int32(40)))) = v468
	*(*int64)(unsafe.Add(mBase, uint32(v470)+8)) = v471
	v478 = int32(24)
	v479 = v66 + v478
	*(*int32)(unsafe.Add(mBase, uint32(v66)+28)) = v479
	v484 = v66 - v478
	*(*int32)(unsafe.Add(mBase, uint32(v66-int32(20)))) = v484
	v489 = v66 - int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v66-int32(28)))) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v484
	*(*int32)(unsafe.Add(mBase, uint32(v489))) = v489
	v495 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	v496 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v495
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v499
	v502 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+48)) = v502
	v505 = *(*int32)(unsafe.Add(mBase, _consts[743]))
	v507 = v66 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v496
	*(*int64)(unsafe.Add(mBase, uint32(v66-int32(8)))) = int64(0)
	v517 = v66 - int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v66-int32(12)))) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v66)+52)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v517
	v522 = int32(*(*uint8)(unsafe.Add(mBase, _consts[809])))
	if v522 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v435)+24))
	if v440 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v443 = v66 + int32(4)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	if v448 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v435
	goto L101
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443))) = v435
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v454)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v66
	v459 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v459+int32(3584))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	goto L95
L103:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v664)+16))
	if v678 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = int32(32)
	v528 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+12))
	if v529 == int32(0) {
		v614 = v528
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+24)) = v656 + int32(1)
	v664 = v655
	goto L103
L107:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	if v627 != v479 {
		goto L124
	} else {
		goto L125
	}
L108:
	;
	v533 = v528 + int32(8)
	if v529 == v533 {
		v614 = v528
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v540 = v529
	goto L110
L110:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+44)))
	if v552&int32(41) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v609 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v614 = v609
	goto L107
L112:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[810]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+4))
	if v559 == int32(0) {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v606 != v533 {
		v540 = v606
		goto L110
	} else {
		goto L123
	}
L115:
	;
	if v559 == v558 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v563)+4)) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v559)+20)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v559)+16)) = v540 + int32(-64)
	v573 = v540 + int32(24)
	v575 = v540 + int32(28)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	if v576 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575))) = v573
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v573
	goto L119
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+4)) = v573
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v582)+4)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v559
	v587 = v66 + int32(28)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if v588 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v479
	v593 = v66 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v593
	goto L122
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+12)) = v479
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+8)) = v597
	v600 = v559 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v597)+4)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v600
	goto L114
L123:
	;
	goto L111
L124:
	;
	v630 = v627
	goto L126
L125:
	;
	v630 = int32(0)
	goto L126
L126:
	;
	if v630 != 0 {
		v664 = v614
		goto L103
	} else {
		goto L127
	}
L127:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v633 = v66 + int32(4)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	*(*int32)(unsafe.Add(mBase, uint32(v631)+4)) = v634
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	if v638 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v614)+4)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v614
	goto L130
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633))) = v614
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v644)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v66
	v649 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v649+int32(3584))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	goto L95
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[811])) = v80
	v756 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[812])) = uint8(v756)
	v759 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v759+int32(3584))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L13
	} else {
		goto L155
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v664)+16)) = v677
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v686 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v690 = F_LWLockAcquire(m, v686+int32(6656), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L13
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	if v677 != v678 {
		goto L132
	} else {
		goto L154
	}
L136:
	;
	if v684 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v740+int32(6656))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L153
	}
L138:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	*(*int64)(unsafe.Add(mBase, uint32(v695)+8)) = int64(0)
	goto L137
L139:
	;
	goto L140
L140:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v700 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	if v710 == int32(0) {
		v734 = v712
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+316))
	v708 = base.B2i32(v706 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v708)
	v710 = v708
	goto L144
L143:
	;
	v710 = int32(0)
	goto L144
L144:
	;
	goto L141
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v734)+12)) = v684
	goto L137
L146:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	if v715 == int32(0) {
		v734 = v712
		goto L145
	} else {
		goto L147
	}
L147:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v715))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v684)) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v729 == int32(0) {
		goto L137
	} else {
		goto L152
	}
L149:
	;
	v729 = base.B2i32(base.Ui32(v684) < base.Ui32(v715))
	goto L148
L150:
	;
	goto L151
L151:
	;
	v729 = int32(base.Ui32(v684-v715) >> (uint(int32(31)) % 32))
	goto L148
L152:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v734 = v733
	goto L145
L153:
	;
	goto L132
L154:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v664)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v664)+20)) = v746 + int32(1)
	goto L132
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(103079215120)
	v769 = *(*int32)(unsafe.Add(mBase, _consts[813]))
	v773 = F_hash_create(m, int32(301847), v769, v18+int32(-48), int32(40))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _consts[814])) = v773
	goto L95
L157:
	;
	F_errmsg_internal(m, int32(247194), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(475280), int32(1784), int32(90986))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+4)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(v816))) = v816
	goto L162
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+68)) = v816
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+64)) = v823
	v826 = v80 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v823)+4)) = v826
	*(*int32)(unsafe.Add(mBase, uint32(v816))) = v826
	v830 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v830+int32(3584))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(82402), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	F_errdetail(m, int32(592262), v20)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L13
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(475280), int32(1829), int32(90986))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L13
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	F_errmsg(m, int32(101898), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L13
	} else {
		goto L171
	}
L171:
	;
	F_errhint(m, int32(622267), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L13
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(475280), int32(679), int32(102176))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L13
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSnapshotData(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
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
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v452 int64
	_ = v452
	var v457 int64
	_ = v457
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v470 int64
	_ = v470
	var v474 int32
	_ = v474
	var v475 int64
	_ = v475
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v491 int64
	_ = v491
	var v499 int64
	_ = v499
	var v503 int64
	_ = v503
	var v506 int64
	_ = v506
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v516 int64
	_ = v516
	var v518 int64
	_ = v518
	var v519 int64
	_ = v519
	var v521 int32
	_ = v521
	var v523 int64
	_ = v523
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v529 int64
	_ = v529
	var v532 int64
	_ = v532
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v538 int64
	_ = v538
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	v2 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v31 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L8
	} else {
		goto L179
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L8
	} else {
		goto L175
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v37 = F_emscripten_builtin_malloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
	if v37 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v58 = F_LWLockAcquire(m, v54+int32(512), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[786]))
	v48 = F_emscripten_builtin_malloc(m, (v42+v44)*int32(260))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v48
	if v48 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(0)
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)+56))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	if v65 == int64(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v30+v97<<(uint(int32(2))%32))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v63)+48))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v106 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	if v65 != v64 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	if v72 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[787])) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v69
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[784])) = v69
	v81 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v81
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v86)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v89+int32(512))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	return l0
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v116)
	v118 = int32(3)
	v119 = base.I32_wrap_i64(v103)
	v121 = v119 + int32(1)
	if base.Ui32(v121) <= base.Ui32(v118) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+316))
	v114 = base.B2i32(v112 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v114)
	v116 = v114
	goto L21
L20:
	;
	v116 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v124 = v118
	goto L24
L23:
	;
	v124 = v121
	goto L24
L24:
	;
	if v101-v124 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v128 = v101
	goto L27
L26:
	;
	v128 = v124
	goto L27
L27:
	;
	if base.Ui32(int32(2)) < base.Ui32(v101) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v131 = v128
	goto L30
L29:
	;
	v131 = v124
	goto L30
L30:
	;
	if v116 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+32))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+28))
	v393 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+40))
	if v394 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v134 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+20))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v236)+16))
	if v237 <= v239 {
		goto L61
	} else {
		goto L62
	}
L35:
	;
	v365 = int32(0)
	v367 = v131
	v370 = v2
	v371 = v2
	goto L31
L36:
	;
	goto L37
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v142 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	v145 = int32(0)
	v148 = v145
	v149 = v145
	v151 = v131
	v154 = v2
	v155 = v2
	goto L38
L38:
	;
	v173 = v148 << (uint(int32(2)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v30+v173)))
	if v175 == int32(0) {
		v227 = v149
		v228 = v151
		v230 = v154
		v231 = v155
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v365 = v227
	v367 = v228
	v370 = v230
	v371 = v231
	goto L31
L40:
	;
	v233 = v148 + int32(1)
	if v233 != v134 {
		v148 = v233
		v149 = v227
		v151 = v228
		v154 = v230
		v155 = v231
		goto L38
	} else {
		goto L59
	}
L41:
	;
	if v148 == v97 {
		v227 = v149
		v228 = v151
		v230 = v154
		v231 = v155
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(0) <= v175-v124 {
		v227 = v149
		v228 = v151
		v230 = v154
		v231 = v155
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v143))))
	if v183&int32(18) != 0 {
		v227 = v149
		v228 = v151
		v230 = v154
		v231 = v155
		goto L40
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140+v155<<(uint(int32(2))%32)))) = v175
	if v175-v151 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v193 = v175
	goto L47
L46:
	;
	v193 = v151
	goto L47
L47:
	;
	v194 = int32(1)
	v195 = v155 + v194
	if v154&v194 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v227 = v149
	v228 = v193
	v230 = int32(1)
	v231 = v195
	goto L40
L49:
	;
	goto L50
L50:
	;
	v199 = int32(1)
	v202 = v144 + v148<<(uint(v199)%32)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v203 != 0 {
		v227 = v149
		v228 = v193
		v230 = v199
		v231 = v195
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v204 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v227 = v149
	v228 = v193
	v230 = int32(0)
	v231 = v195
	goto L40
L53:
	;
	goto L54
L54:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v210 = int32(2)
	v214 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v173+(v27+int32(36)))))
	v223 = v204 << (uint(v210) % 32)
	if v223 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v227 = v149 + v204
	v228 = v193
	v230 = int32(0)
	v231 = v195
	goto L40
L56:
	;
	v224 = F__emscripten_memcpy_bulkmem(m, v209+v149<<(uint(v210)%32), v214+v216*int32(640)+int32(280), v223)
	mBase = m.M
	goto L58
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L39
L60:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v350))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v327)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	v325 = int32(0)
	v327 = v131
	goto L60
L62:
	;
	goto L63
L63:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	v246 = v239
	v247 = int32(0)
	v249 = v131
	v250 = v243
	goto L64
L64:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v250))))
	if v271 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v325 = v316
	v327 = v318
	goto L60
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _consts[771]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v246<<(uint(int32(2))%32))))
	if v247 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v316 = v247
	v318 = v249
	v319 = v250
	goto L68
L68:
	;
	v321 = v246 + int32(1)
	if v321 != v237 {
		v246 = v321
		v247 = v316
		v249 = v318
		v250 = v319
		goto L64
	} else {
		goto L84
	}
L69:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v249))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v279)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v295 = v249
	goto L71
L71:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v279)) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	if v293 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v293 = base.B2i32(base.Ui32(v279) < base.Ui32(v249))
	goto L72
L74:
	;
	goto L75
L75:
	;
	v293 = int32(base.Ui32(v279-v249) >> (uint(int32(31)) % 32))
	goto L72
L76:
	;
	v294 = v279
	goto L78
L77:
	;
	v294 = v249
	goto L78
L78:
	;
	v295 = v294
	goto L71
L79:
	;
	if v307 != 0 {
		v325 = v247
		v327 = v295
		goto L60
	} else {
		goto L83
	}
L80:
	;
	v307 = base.B2i32(base.Ui32(v124) <= base.Ui32(v279))
	goto L79
L81:
	;
	goto L82
L82:
	;
	v307 = base.B2i32(int32(0) <= v279-v124)
	goto L79
L83:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[770]))
	*(*int32)(unsafe.Add(mBase, uint32(v238+v247<<(uint(int32(2))%32)))) = v279
	v316 = v247 + int32(1)
	v318 = v295
	v319 = v309
	goto L68
L84:
	;
	goto L65
L85:
	;
	v365 = v325
	v367 = v327
	v370 = v362
	v371 = v2
	goto L31
L86:
	;
	v362 = base.B2i32(base.Ui32(v327) <= base.Ui32(v350))
	goto L85
L87:
	;
	goto L88
L88:
	;
	v362 = base.B2i32(v327-v350 <= int32(0))
	goto L85
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[787])) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v393)+40)) = v367
	goto L91
L90:
	;
	goto L91
L91:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v401+int32(512))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	if v367 == int32(0) {
		v424 = v391
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v445 = v103 + base.I64_extend_i32_s(v424-v119)
	v447 = *(*int64)(unsafe.Add(mBase, _consts[780]))
	v452 = v103 + base.I64_extend_i32_s(v442-v119)
	if base.I32_wrap_i64(v452) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L94:
	;
	if v390 == int32(0) {
		v442 = v424
		goto L93
	} else {
		goto L104
	}
L95:
	;
	if v391 == int32(0) {
		v424 = v367
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v391))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v367)) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v421 != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v421 = base.B2i32(base.Ui32(v367) < base.Ui32(v391))
	goto L97
L99:
	;
	goto L100
L100:
	;
	v421 = int32(base.Ui32(v367-v391) >> (uint(int32(31)) % 32))
	goto L97
L101:
	;
	v422 = v367
	goto L103
L102:
	;
	v422 = v391
	goto L103
L103:
	;
	v424 = v422
	goto L94
L104:
	;
	if v424 == int32(0) {
		v442 = v390
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v424))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v390)) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v440 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v440 = base.B2i32(base.Ui32(v390) < base.Ui32(v424))
	goto L106
L108:
	;
	goto L109
L109:
	;
	v440 = int32(base.Ui32(v390-v424) >> (uint(int32(31)) % 32))
	goto L106
L110:
	;
	v441 = v390
	goto L112
L111:
	;
	v441 = v424
	goto L112
L112:
	;
	v442 = v441
	goto L93
L113:
	;
	*(*int64)(unsafe.Add(mBase, _consts[781])) = v470
	*(*int64)(unsafe.Add(mBase, _consts[780])) = v468
	v474 = int32(4359272)
	v475 = *(*int64)(unsafe.Add(mBase, _consts[782]))
	if base.Ui64(v475) < base.Ui64(v445) {
		goto L129
	} else {
		goto L130
	}
L114:
	;
	v457 = *(*int64)(unsafe.Add(mBase, _consts[781]))
	v468 = v447
	v470 = v457
	goto L113
L115:
	;
	goto L116
L116:
	;
	if base.Ui64(v447) < base.Ui64(v452) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v459 = v452
	goto L119
L118:
	;
	v459 = v447
	goto L119
L119:
	;
	if base.I32_wrap_i64(v447) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v461 = v459
	goto L122
L121:
	;
	v461 = v452
	goto L122
L122:
	;
	v463 = *(*int64)(unsafe.Add(mBase, _consts[781]))
	if base.Ui64(v463) < base.Ui64(v452) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v465 = v452
	goto L125
L124:
	;
	v465 = v463
	goto L125
L125:
	;
	if base.I32_wrap_i64(v463) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v467 = v465
	goto L128
L127:
	;
	v467 = v452
	goto L128
L128:
	;
	v468 = v461
	v470 = v467
	goto L113
L129:
	;
	v479 = v445
	goto L131
L130:
	;
	v479 = v475
	goto L131
L131:
	;
	if base.I32_wrap_i64(v475) != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v481 = v479
	goto L134
L133:
	;
	v481 = v445
	goto L134
L134:
	;
	if base.I32_wrap_i64(v445) != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v483 = v481
	goto L137
L136:
	;
	v483 = v475
	goto L137
L137:
	;
	*(*int64)(unsafe.Add(mBase, _consts[782])) = v483
	if base.Ui32(int32(3)) <= base.Ui32(v101) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v508 = v103 + base.I64_extend_i32_s(v101-v119)
	goto L140
L139:
	;
	v491 = int64(1)
	v499 = v103 + v491
	if base.Ui32(base.I32_wrap_i64(v499)) < base.Ui32(int32(3)) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	*(*int64)(unsafe.Add(mBase, _consts[783])) = v508
	v510 = v103 + base.I64_extend_i32_s(v102-v119)
	v511 = int32(4359248)
	v513 = *(*int64)(unsafe.Add(mBase, _consts[776]))
	if base.I32_wrap_i64(v513) != 0 {
		goto L147
	} else {
		goto L148
	}
L141:
	;
	v503 = v103 + (v491-v103)&int64(4294967295) + int64(2)
	goto L143
L142:
	;
	v503 = v499
	goto L143
L143:
	;
	if base.Ui64(int64(2)) < base.Ui64(v499) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v506 = v503
	goto L146
L145:
	;
	v506 = v499
	goto L146
L146:
	;
	v508 = v506
	goto L140
L147:
	;
	if base.Ui64(v510) < base.Ui64(v513) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v519 = v510
	goto L149
L149:
	;
	*(*int64)(unsafe.Add(mBase, _consts[776])) = v519
	v521 = int32(4359264)
	v523 = *(*int64)(unsafe.Add(mBase, _consts[777]))
	if base.I32_wrap_i64(v523) != 0 {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	v516 = v513
	goto L152
L151:
	;
	v516 = v510
	goto L152
L152:
	;
	if base.I32_wrap_i64(v510) != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v518 = v516
	goto L155
L154:
	;
	v518 = v513
	goto L155
L155:
	;
	v519 = v518
	goto L149
L156:
	;
	if base.Ui64(v510) < base.Ui64(v523) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v529 = v510
	goto L158
L158:
	;
	*(*int64)(unsafe.Add(mBase, _consts[777])) = v529
	v532 = *(*int64)(unsafe.Add(mBase, _consts[778]))
	if base.I32_wrap_i64(v532) != 0 {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	v526 = v523
	goto L161
L160:
	;
	v526 = v510
	goto L161
L161:
	;
	if base.I32_wrap_i64(v510) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v528 = v526
	goto L164
L163:
	;
	v528 = v523
	goto L164
L164:
	;
	v529 = v528
	goto L158
L165:
	;
	if base.Ui64(v510) < base.Ui64(v532) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	v538 = v510
	goto L167
L167:
	;
	*(*int64)(unsafe.Add(mBase, _consts[779])) = v508
	*(*int64)(unsafe.Add(mBase, _consts[778])) = v538
	*(*int32)(unsafe.Add(mBase, _consts[784])) = v367
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v64
	v547 = v370 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v547)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v554 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L8
	} else {
		goto L174
	}
L168:
	;
	v535 = v532
	goto L170
L169:
	;
	v535 = v510
	goto L170
L170:
	;
	if base.I32_wrap_i64(v510) != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v537 = v535
	goto L173
L172:
	;
	v537 = v532
	goto L173
L173:
	;
	v538 = v537
	goto L167
L174:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v554
	v559 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v559)
	return l0
L175:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L8
	} else {
		goto L176
	}
L176:
	;
	F_errmsg(m, int32(12790), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(469465), int32(2217), int32(481624))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L8
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L8
	} else {
		goto L180
	}
L180:
	;
	F_errmsg(m, int32(12790), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L8
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(469465), int32(2224), int32(481624))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L8
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetVisibilityMapPins(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	v8 = int32(0)
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v27 = v21 ^ int32(-1)
	v29 = v21 << (uint(int32(13)) % 32)
	if int32(0) <= v21 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v20 = l3
	v21 = l1
	v22 = l5
	v23 = l6
	v24 = l4
	v25 = l2
	goto L1
L3:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v20 = l4
	v21 = l2
	v22 = l6
	v23 = l5
	v24 = l3
	v25 = l1
	goto L1
L6:
	;
	if base.Ui32(l3) <= base.Ui32(l4) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return v256
L9:
	;
	F_LockBuffer(m, v21, int32(2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L39
	} else {
		goto L54
	}
L10:
	;
	F_visibilitymap_pin(m, l0, v24, v23)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L39
	} else {
		goto L53
	}
L11:
	;
	if v95 != 0 {
		goto L10
	} else {
		goto L52
	}
L12:
	;
	if v55 != 0 {
		v256 = v8
		goto L8
	} else {
		goto L48
	}
L13:
	;
	v64 = v25 ^ int32(-1)
	v66 = v25 << (uint(int32(13)) % 32)
	if int32(0) <= v25 {
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+10)))
	if v44&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v43 = v33 + v29 + int32(-8192)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v27<<(uint(int32(2))%32))))
	v43 = v42
	goto L14
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v47 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	if v25 == int32(0) {
		v256 = v8
		goto L8
	} else {
		goto L26
	}
L21:
	;
	if v25 == int32(0) {
		goto L12
	} else {
		goto L25
	}
L22:
	;
	v55 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v51 = F_BufferGetBlockNumber(m, v47)
	mBase = m.M
	v53 = base.I32_div_u_s(v20, int32(32672))
	v55 = base.B2i32(v51 == v53)
	goto L21
L25:
	;
	v62 = v55 ^ int32(1)
	goto L13
L26:
	;
	v62 = v8
	goto L13
L27:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
	if v81&int32(4) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v80 = v70 + v66 + int32(-8192)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v64<<(uint(int32(2))%32))))
	v80 = v79
	goto L27
L31:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v84 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v95 = v8
	goto L33
L33:
	;
	if (v62|v95)&int32(1) == int32(0) {
		v256 = v8
		goto L8
	} else {
		goto L38
	}
L34:
	;
	v95 = v92 ^ int32(1)
	goto L33
L35:
	;
	v92 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v88 = F_BufferGetBlockNumber(m, v84)
	mBase = m.M
	v90 = base.I32_div_u_s(v24, int32(32672))
	v92 = base.B2i32(v88 == v90)
	goto L34
L38:
	;
	F_LockBuffer(m, v21, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	v106 = int32(0)
	v109 = base.B2i32(v25 == v106) | base.B2i32(l1 == l2)
	if v109 == v106 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_LockBuffer(m, v25, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v62 == int32(0) {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	F_visibilitymap_pin(m, l0, v20, v22)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	if v95 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v136 = int32(1)
	v137 = int32(0)
	goto L9
L48:
	;
	F_LockBuffer(m, v21, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	F_visibilitymap_pin(m, l0, v20, v22)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	F_LockBuffer(m, v21, int32(2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	return int32(1)
L52:
	;
	v131 = int32(0)
	v136 = v131
	v137 = v131
	goto L9
L53:
	;
	v136 = v62
	v137 = int32(1)
	goto L9
L54:
	;
	v141 = int32(1)
	if v109 != 0 {
		v256 = v141
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	if v137&v136 != 0 {
		v256 = v141
		goto L8
	} else {
		goto L57
	}
L57:
	;
	goto L58
L58:
	;
	v166 = int32(0)
	if base.B2i32(int32(0) <= v21) == v166 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v256 = v141
	goto L8
L60:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
	if v179&int32(4) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170+v27<<(uint(int32(2))%32))))
	v178 = v172
	goto L60
L62:
	;
	goto L63
L63:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v178 = v174 + v29 + int32(-8192)
	goto L60
L64:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v182 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v193 = v166
	goto L66
L66:
	;
	v194 = int32(0)
	if v25 < v194 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v193 = v190 ^ int32(1)
	goto L66
L68:
	;
	v190 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v186 = F_BufferGetBlockNumber(m, v182)
	mBase = m.M
	v188 = base.I32_div_u_s(v20, int32(32672))
	v190 = base.B2i32(v186 == v188)
	goto L67
L71:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
	if v209&int32(4) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+v64<<(uint(int32(2))%32))))
	v208 = v202
	goto L71
L73:
	;
	goto L74
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v208 = v204 + v66 + int32(-8192)
	goto L71
L75:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v212 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v223 = v194
	goto L77
L77:
	;
	if (v193|v223)&int32(1) == int32(0) {
		v256 = v141
		goto L8
	} else {
		goto L82
	}
L78:
	;
	v223 = v220 ^ int32(1)
	goto L77
L79:
	;
	v220 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v216 = F_BufferGetBlockNumber(m, v212)
	mBase = m.M
	v218 = base.I32_div_u_s(v24, int32(32672))
	v220 = base.B2i32(v216 == v218)
	goto L78
L82:
	;
	F_LockBuffer(m, v21, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L39
	} else {
		goto L83
	}
L83:
	;
	F_LockBuffer(m, v25, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L39
	} else {
		goto L84
	}
L84:
	;
	if v193 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_visibilitymap_pin(m, l0, v20, v22)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L39
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if v223 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	F_visibilitymap_pin(m, l0, v24, v23)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_LockBuffer(m, v21, int32(2))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L39
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L39
	} else {
		goto L94
	}
L94:
	;
	if v193&v223 != int32(1) {
		goto L58
	} else {
		goto L95
	}
L95:
	;
	goto L59
}
func F___getf2(m *base.Module, l0 int64, l1 int64, l2 int64) int32 {
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = int32(-1)
	v11 = l1 & int64(9223372036854775807)
	v12 = int64(9223090561878065152)
	if v11 == v12 {
		v16 = base.B2i32(l0 != int64(0))
	} else {
		v16 = base.B2i32(base.Ui64(v12) < base.Ui64(v11))
	}
	if v16 != 0 {
		v50 = v7
		return v50
	} else {
		v18 = l2 & int64(9223372036854775807)
		v19 = int64(9223090561878065152)
		if base.B2i32(base.Ui64(v19) < base.Ui64(v18))&base.B2i32(v18 != v19) != 0 {
			v50 = v7
			return v50
		} else {
			if l0|(v11|v18) == int64(0) {
				return int32(0)
			} else {
				if int64(0) <= l1&l2 {
					if base.B2i32(l1 != l2)&base.B2i32(l1 < l2) != 0 {
						v50 = v7
						return v50
					} else {
						return base.B2i32(l0|(l1^l2) != int64(0))
					}
				} else {
					if l1 == l2 {
						v45 = base.B2i32(l0 != int64(0))
					} else {
						v45 = base.B2i32(l2 < l1)
					}
					if v45 != 0 {
						v50 = v7
					} else {
						v50 = base.B2i32(l0|(l1^l2) != int64(0))
					}
					return v50
				}
			}
		}
	}
}
func F_gbk_to_utf8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(37), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4330660), v18, v18, v18, int32(37), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_generateClonedIndexStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	v26 = m.G0
	v28 = v26 - int32(144)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v34 = F_SearchSysCache1(m, int32(57), v30)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v689 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+8)))
	if v667 < v689 {
		goto L158
	} else {
		goto L159
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L12
	} else {
		goto L153
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L12
	} else {
		goto L150
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L12
	} else {
		goto L145
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L12
	} else {
		goto L142
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L12
	} else {
		goto L139
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L12
	} else {
		goto L136
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L12
	} else {
		goto L133
	}
L12:
	;
	return int32(0)
L13:
	;
	if v34 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v48 = F_SearchSysCache1(m, int32(2), v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L12
	} else {
		goto L130
	}
L17:
	;
	if v48 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v56 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(17))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v60 = F_SysCacheGetAttrNotNull(m, int32(34), v38, int32(18))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v63 = F_palloc0(m, int32(72))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(204)
	v71 = F_pstrdup(m, v52+v53+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = F_get_tablespace_name(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	v78 = int32(0)
	goto L25
L25:
	;
	v79 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v63)+36)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v63)+44)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v63)+52)) = v79
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+60)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+61)) = uint8(v88)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+62)) = uint8(v90)
	if v90 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v78 = v75
	goto L25
L27:
	;
	v100 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+67)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = int32(0)
	v105 = v99 & v100
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+64)) = uint8(v105)
	if v90 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+12)))
	if v95 != int32(1) {
		v99 = int32(0)
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	v99 = v98
	goto L27
L31:
	;
	goto L30
L32:
	;
	v282 = F_SysCacheGetAttr(m, int32(34), v38, int32(20), v28+int32(135))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L12
	} else {
		goto L66
	}
L33:
	;
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v251)
	goto L32
L34:
	;
	v112 = F_get_index_constraint(m, v30)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	if v86&int32(1) != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	if v109 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	if v112 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l3 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v249 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v249)
	goto L32
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112
	goto L44
L43:
	;
	goto L44
L44:
	;
	v116 = F_SearchSysCache1(m, int32(19), v112)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	if v116 == int32(0) {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+63)) = uint8(v122)
	v124 = v121 + v120
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+65)) = uint8(v125)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+66)) = uint8(v127)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)))
	if v129 != v122 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ReleaseCatCache(m, v116)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L12
	} else {
		goto L65
	}
L48:
	;
	v135 = F_SysCacheGetAttrNotNull(m, int32(19), v116, int32(27))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v137 = F_pg_detoast_datum(m, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_deconstruct_array_builtin(m, v137, int32(26), v28+int32(140), int32(0), v28+int32(136))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v28)+136))
	if v147 <= int32(0) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v151 = v63 + int32(36)
	v155 = int32(0)
	goto L53
L53:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v155<<(uint(int32(2))%32))))
	v183 = F_SearchSysCache1(m, int32(40), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L55
	}
L54:
	;
	goto L47
L55:
	;
	if v183 == int32(0) {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v192 = F_pstrdup(m, v189+int32(4))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+68))
	v195 = F_get_namespace_name(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	v197 = F_makeString(m, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v197
	v200 = F_makeString(m, v192)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v28)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v204
	v210 = F_list_make2_impl(m, v28+int32(120), v28+int32(116))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v213 = F_lappend(m, v212, v210)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v213
	F_ReleaseCatCache(m, v183)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v219 = v155 + int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v28)+136))
	if v219 < v220 {
		v155 = v219
		goto L53
	} else {
		goto L64
	}
L64:
	;
	goto L54
L65:
	;
	goto L32
L66:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v284 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v313 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+10)))
	if v313 <= int32(0) {
		v667 = v313
		goto L4
	} else {
		goto L76
	}
L68:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
	v309 = v302
	v310 = v295
	v311 = v300
	v312 = v308
	goto L67
L69:
	;
	v306 = int32(0)
	v309 = v303
	v310 = v306
	v311 = v305
	v312 = v306
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v63)+20)) = int64(0)
	v303 = v63 + int32(20)
	v305 = v63 + int32(24)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v293 = F_text_to_cstring(m, v282)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v295 = F_stringToNode(m, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v63)+20)) = int64(0)
	v300 = v63 + int32(24)
	v302 = v63 + int32(20)
	if v295 != 0 {
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v303 = v302
	v305 = v300
	goto L69
L76:
	;
	v316 = int32(24)
	v323 = int32(0)
	v336 = v312
	goto L77
L77:
	;
	v349 = v323 << (uint(int32(1)) % 32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349+v350))))
	v354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349+(v41+int32(48))))))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v358 = F_palloc0(m, int32(36))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L12
	} else {
		goto L79
	}
L78:
	;
	v667 = v535
	goto L4
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = int32(92)
	if v354 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v409 = F_pstrdup(m, v355+v356<<(uint(int32(4))%32)+v323*int32(100)+int32(24))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L12
	} else {
		goto L93
	}
L81:
	;
	v363 = F_get_attname(m, v42, v354, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v336 == int32(0) {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	v365 = F_get_atttype(m, v42, v354)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+4)) = v363
	v398 = v365
	v400 = v336
	goto L80
L86:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v379 = F_map_variable_attnos(m, v374, int32(1), l2, int32(0), v28+int32(140))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+140)))
	if v381 == int32(1) {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+8)) = v379
	v385 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+4)) = v385
	v388 = v336 + int32(4)
	if base.Ui32(v388) < base.Ui32(v373+v372<<(uint(int32(2))%32)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v394 = v388
	goto L91
L90:
	;
	v394 = v385
	goto L91
L91:
	;
	v395 = F_exprType(m, v379)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L12
	} else {
		goto L92
	}
L92:
	;
	v398 = v395
	v400 = v394
	goto L80
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = v409
	v412 = int32(0)
	v414 = v323 << (uint(int32(2)) % 32)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v56+v316+v414)))
	if v416 == v412 {
		v454 = v412
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+16)) = v454
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v414+(v60+v316))))
	v461 = F_SearchSysCache1(m, int32(14), v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L12
	} else {
		goto L106
	}
L95:
	;
	v419 = F_get_typcollation(m, v398)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	if v419 == v416 {
		v454 = v412
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v423 = F_SearchSysCache1(m, int32(16), v416)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L12
	} else {
		goto L98
	}
L98:
	;
	if v423 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423)+16))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+22)))
	v429 = v427 + v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+68))
	v431 = F_get_namespace_name(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v435 = F_pstrdup(m, v429+int32(4))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	v437 = F_makeString(m, v431)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L12
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v437
	v440 = F_makeString(m, v435)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+136)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v440
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v444
	v450 = F_list_make2_impl(m, v28+int32(92), v28+int32(88))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	F_ReleaseCatCache(m, v423)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	v454 = v450
	goto L94
L106:
	;
	if v461 == int32(0) {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v461)+16))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466)+22)))
	v468 = v466 + v467
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v470 = F_GetDefaultOpClass(m, v398, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	if v470 != v460 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v468)+72))
	v474 = F_get_namespace_name(m, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	v496 = int32(0)
	goto L111
L111:
	;
	F_ReleaseCatCache(m, v461)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L12
	} else {
		goto L117
	}
L112:
	;
	v478 = F_pstrdup(m, v468+int32(8))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v480 = F_makeString(m, v474)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v480
	v483 = F_makeString(m, v478)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+136)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v483
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v487
	v493 = F_list_make2_impl(m, v28+int32(76), v28+int32(72))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	v496 = v493
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+20)) = v496
	v501 = v323 + int32(1)
	v503 = F_get_attoptions(m, v30, base.I32_extend16_s(v501))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	v505 = F_untransformRelOptions(m, v503)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v358)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+24)) = v505
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+10)))
	if v511 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v532 = F_lappend(m, v531, v358)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L12
	} else {
		goto L128
	}
L121:
	;
	if v352&int32(1) != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358)+32)) = v528
	goto L120
L123:
	;
	v516 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+28)) = v516
	if v352&v516 == int32(0) {
		v528 = v516
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v352&int32(2) == int32(0) {
		goto L120
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v528 = int32(1)
	goto L122
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v532
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+10)))
	if v501 < v535 {
		v323 = v501
		v336 = v400
		goto L77
	} else {
		goto L129
	}
L129:
	;
	goto L78
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	F_errmsg_internal(m, int32(43821), v28)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(476263), int32(1723), int32(91144))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v554
	F_errmsg_internal(m, int32(50732), v28+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(476263), int32(1735), int32(91144))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v112
	F_errmsg_internal(m, int32(38585), v28+int32(96))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(476263), int32(1800), int32(91144))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L12
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v182
	F_errmsg_internal(m, int32(40645), v28+int32(112))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L12
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(476263), int32(1835), int32(91144))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L12
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errmsg_internal(m, int32(71138), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(476263), int32(1902), int32(91144))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L12
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(397140), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v620 + int32(4)
	F_errdetail(m, int32(602995), v28+int32(32))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(476263), int32(1918), int32(91144))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L12
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v416
	F_errmsg_internal(m, int32(43493), v28+int32(80))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L12
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(476263), int32(2188), int32(249828))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L12
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v460
	F_errmsg_internal(m, int32(39836), v28+int32(48))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(476263), int32(2215), int32(121237))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L12
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L12
	} else {
		goto L188
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L12
	} else {
		goto L184
	}
L158:
	;
	v696 = v667
	goto L161
L159:
	;
	goto L160
L160:
	;
	v785 = F_SysCacheGetAttr(m, int32(57), v34, int32(33), v28+int32(135))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L12
	} else {
		goto L169
	}
L161:
	;
	v721 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+int32(48)+v696<<(uint(int32(1))%32)))))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	v725 = F_palloc0(m, int32(36))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L12
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = int32(92)
	if v721 == int32(0) {
		goto L157
	} else {
		goto L164
	}
L164:
	;
	v732 = F_get_attname(m, v42, v721, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L12
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v725)+4)) = v732
	v745 = F_pstrdup(m, v722+v723<<(uint(int32(4))%32)+v696*int32(100)+int32(24))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L12
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v725)+12)) = v745
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v749 = F_lappend(m, v748, v725)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L12
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v749
	v753 = v696 + int32(1)
	v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+8)))
	if v753 < v754 {
		v696 = v753
		goto L161
	} else {
		goto L168
	}
L168:
	;
	goto L162
L169:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v787 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v790 = F_untransformRelOptions(m, v785)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L12
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v797 = F_SysCacheGetAttr(m, int32(34), v38, int32(21), v28+int32(135))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L12
	} else {
		goto L174
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+28)) = v790
	goto L172
L174:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+135)))
	if v799 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v802 = F_text_to_cstring(m, v797)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L12
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	F_ReleaseCatCache(m, v34)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L12
	} else {
		goto L182
	}
L178:
	;
	v804 = F_stringToNode(m, v802)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	v810 = F_map_variable_attnos(m, v804, int32(1), l2, int32(0), v28+int32(140))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+140)))
	if v812 == int32(1) {
		goto L156
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+32)) = v810
	goto L177
L182:
	;
	F_ReleaseCatCache(m, v48)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L12
	} else {
		goto L183
	}
L183:
	;
	m.G0 = v28 + int32(144)
	return v63
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(139076), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(476263), int32(1988), int32(91144))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(397140), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v852 + int32(4)
	F_errdetail(m, int32(602995), v28-int32(-64))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(476263), int32(2026), int32(91144))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_generate_combinations_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 < v6 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 <= l2 {
		} else {
			v11 = l1 + int32(1)
			v17 = l2
			for {
				*(*int32)(unsafe.Add(mBase, uint32(l3+l1<<(uint(int32(2))%32)))) = v17
				v22 = v17 + int32(1)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v11 < v24 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v26 <= v22 {
					} else {
						v83 = int32(2)
						v35 = v22
						for {
							*(*int32)(unsafe.Add(mBase, uint32(l3+v11<<(uint(v83)%32)))) = v35
							v40 = v35 + int32(1)
							F_generate_combinations_recurse(m, l0, l1+v83, v40, l3)
							mBase = m.M
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v40 < v42 {
								v35 = v40
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v47 = int32(2)
					v52 = F___memcpy(m, v44+v45*v24<<(uint(v47)%32), l3, v24<<(uint(v47)%32))
					mBase = m.M
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v53 + int32(1)
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v22 < v62 {
					v17 = v22
					continue
				} else {
					break
				}
				break
			}
		}
	} else {
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v67 = int32(2)
		v71 = v6 << (uint(v67) % 32)
		if v71 != 0 {
			v72 = F__emscripten_memcpy_bulkmem(m, v64+v65*v6<<(uint(v67)%32), l3, v71)
			mBase = m.M
		} else {
		}
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v74 + int32(1)
	}
	return
}
func F_generate_implied_equalities_for_column(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v665 int32
	_ = v665
	var v675 int32
	_ = v675
	v6 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = F_find_childrel_parents(m, l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = v6
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = v23
	goto L3
L6:
	;
	return v675
L7:
	;
	if v85 < int32(0) {
		v675 = v6
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v85 = base.I32_ctz(v71) | v72<<(uint(int32(5))%32)
	goto L7
L9:
	;
	v85 = int32(-2)
	goto L7
L10:
	;
	v38 = base.I32_div_s(int32(0), int32(32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 <= v38 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v42 = v28 + int32(8)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32))))
	v49 = v46 & int32(-1)
	if v49 != 0 {
		v71 = v49
		v72 = v38
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = v38 + int32(1)
	if v51 == v39 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v54 = v51
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42+v54<<(uint(int32(2))%32))))
	if v61 != 0 {
		v71 = v61
		v72 = v54
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v63 = v54 + int32(1)
	if v63 != v39 {
		v54 = v63
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v101 = v85
	goto L19
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v101<<(uint(int32(2))%32))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+40)))
	if v113 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v675 = int32(0)
	goto L6
L21:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v609 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L22:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v114 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 < int32(2) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v122 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v123 = v120
	goto L27
L26:
	;
	v123 = int32(0)
	goto L27
L27:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v131 = v124
	v134 = v114
	v135 = int32(-1)
	goto L28
L28:
	;
	if v131 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v322 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L30:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if v252 == int32(0) {
		goto L21
	} else {
		goto L50
	}
L31:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v238 = v131
	v239 = v145
	v241 = v134
	v242 = v135
	goto L30
L32:
	;
	goto L33
L33:
	;
	v155 = v135
	goto L34
L34:
	;
	if v123 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v238 = v232
	v239 = v232
	v241 = v229
	v242 = v220
	goto L30
L36:
	;
	if v220 <= int32(0) {
		goto L21
	} else {
		goto L47
	}
L37:
	;
	v220 = base.I32_ctz(v206) | v207<<(uint(int32(5))%32)
	goto L36
L38:
	;
	v220 = int32(-2)
	goto L36
L39:
	;
	v171 = v155 + int32(1)
	v173 = base.I32_div_s(v171, int32(32))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v174 <= v173 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v177 = v123 + int32(8)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v173<<(uint(int32(2))%32))))
	v184 = v181 & (int32(-1) << (uint(v171) % 32))
	if v184 != 0 {
		v206 = v184
		v207 = v173
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v186 = v173 + int32(1)
	if v186 == v174 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v189 = v186
	goto L43
L43:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v177+v189<<(uint(int32(2))%32))))
	if v196 != 0 {
		v206 = v196
		v207 = v189
		goto L37
	} else {
		goto L45
	}
L44:
	;
	goto L38
L45:
	;
	v198 = v189 + int32(1)
	if v198 != v174 {
		v189 = v198
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v223 <= v220 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v220<<(uint(int32(2))%32))))
	if v229 == int32(0) {
		v155 = v220
		goto L34
	} else {
		goto L49
	}
L49:
	;
	goto L35
L50:
	;
	v256 = v238 + int32(4)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if base.Ui32(v256) < base.Ui32(v239+v258<<(uint(int32(2))%32)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v263 = v256
	goto L53
L52:
	;
	v263 = int32(0)
	goto L53
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v266 = int32(0)
	v273 = base.B2i32(v264|v265 == v266)
	if v264 == v266 {
		v312 = v273
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v312 == int32(0) {
		v131 = v263
		v134 = v241
		v135 = v242
		goto L28
	} else {
		goto L66
	}
L55:
	;
	goto L54
L56:
	;
	if v265 == int32(0) {
		v312 = v273
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v279 != v280 {
		v312 = int32(0)
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v282 = int32(1)
	if v279 <= v282 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v285 = v282
	goto L61
L60:
	;
	v285 = v279
	goto L61
L61:
	;
	v286 = int32(8)
	v291 = int32(0)
	goto L62
L62:
	;
	v299 = v291 << (uint(int32(2)) % 32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v264+v286+v299)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v265+v286))))
	v304 = base.B2i32(v301 == v303)
	if v303 != v301 {
		v312 = v304
		goto L55
	} else {
		goto L64
	}
L63:
	;
	v312 = v304
	goto L55
L64:
	;
	v307 = v291 + int32(1)
	if v307 != v285 {
		v291 = v307
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v318 = m.T0[l2].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, v112, v252, l3)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	if v318 == int32(0) {
		v131 = v263
		v134 = v241
		v135 = v242
		goto L28
	} else {
		goto L68
	}
L68:
	;
	goto L29
L69:
	;
	if v577 != 0 {
		v675 = v577
		goto L6
	} else {
		goto L143
	}
L70:
	;
	v577 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v326 = int32(0)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v328 <= v326 {
		v577 = v326
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v337 = v326
	v340 = v326
	goto L74
L74:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+v340<<(uint(int32(2))%32))))
	if v354 == v252 {
		v554 = v337
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v577 = v554
	goto L69
L76:
	;
	v568 = v340 + int32(1)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v568 < v569 {
		v337 = v554
		v340 = v568
		goto L74
	} else {
		goto L142
	}
L77:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v358 = int32(0)
	if v356 == v358 {
		v399 = v358
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v399 != 0 {
		v554 = v337
		goto L76
	} else {
		goto L92
	}
L79:
	;
	goto L78
L80:
	;
	if v357 == int32(0) {
		v399 = v358
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	if v367 < v368 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v370 = v367
	goto L84
L83:
	;
	v370 = v368
	goto L84
L84:
	;
	if v370 <= int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v373 = int32(1)
	goto L87
L86:
	;
	v373 = v370
	goto L87
L87:
	;
	v374 = int32(8)
	v379 = int32(0)
	goto L88
L88:
	;
	v386 = v379 << (uint(int32(2)) % 32)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v357+v374+v386)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386+(v356+v374))))
	v391 = v388 & v390
	v393 = base.B2i32(v391 != int32(0))
	if v391 != 0 {
		v399 = v393
		goto L79
	} else {
		goto L90
	}
L89:
	;
	v399 = v393
	goto L79
L90:
	;
	v395 = v379 + int32(1)
	if v395 != v373 {
		v379 = v395
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
	v404 = int32(0)
	if v403 == v404 {
		v445 = v404
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v445 != 0 {
		v554 = v337
		goto L76
	} else {
		goto L107
	}
L94:
	;
	goto L93
L95:
	;
	if l4 == int32(0) {
		v445 = v404
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v413 < v414 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v416 = v413
	goto L99
L98:
	;
	v416 = v414
	goto L99
L99:
	;
	if v416 <= int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v419 = int32(1)
	goto L102
L101:
	;
	v419 = v416
	goto L102
L102:
	;
	v420 = int32(8)
	v425 = int32(0)
	goto L103
L103:
	;
	v432 = v425 << (uint(int32(2)) % 32)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l4+v420+v432)))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+(v403+v420))))
	v437 = v434 & v436
	v439 = base.B2i32(v437 != int32(0))
	if v437 != 0 {
		v445 = v439
		goto L94
	} else {
		goto L105
	}
L104:
	;
	v445 = v439
	goto L94
L105:
	;
	v441 = v425 + int32(1)
	if v441 != v419 {
		v425 = v441
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	if v20 == int32(2) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
	v452 = int32(0)
	if v27 == v452 {
		v493 = v452
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v497 == int32(0) {
		v554 = v337
		goto L76
	} else {
		goto L126
	}
L111:
	;
	if v493 != 0 {
		v554 = v337
		goto L76
	} else {
		goto L125
	}
L112:
	;
	goto L111
L113:
	;
	if v451 == int32(0) {
		v493 = v452
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v461 < v462 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v464 = v461
	goto L117
L116:
	;
	v464 = v462
	goto L117
L117:
	;
	if v464 <= int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v467 = int32(1)
	goto L120
L119:
	;
	v467 = v464
	goto L120
L120:
	;
	v468 = int32(8)
	v473 = int32(0)
	goto L121
L121:
	;
	v480 = v473 << (uint(int32(2)) % 32)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v451+v468+v480)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480+(v27+v468))))
	v485 = v482 & v484
	v487 = base.B2i32(v485 != int32(0))
	if v485 != 0 {
		v493 = v487
		goto L112
	} else {
		goto L123
	}
L122:
	;
	v493 = v487
	goto L112
L123:
	;
	v489 = v473 + int32(1)
	if v489 != v467 {
		v473 = v489
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	goto L110
L126:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v500 <= int32(0) {
		v554 = v337
		goto L76
	} else {
		goto L127
	}
L127:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	v511 = int32(0)
	goto L128
L128:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525+v511<<(uint(int32(2))%32))))
	v531 = F_get_opfamily_member_for_cmptype(m, v529, v504, v503, int32(3))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L131
	}
L129:
	;
	v544 = F_create_join_clause(m, l0, v112, v531, v252, v354, v112)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L140
	}
L130:
	;
	goto L129
L131:
	;
	if v531 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v112)+52))
	if v533 == int32(0) {
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v541 = v511 + int32(1)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v541 < v542 {
		v511 = v541
		goto L128
	} else {
		goto L139
	}
L135:
	;
	v536 = F_get_opcode(m, v531)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v538 = F_get_func_leakproof(m, v536)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v538 != 0 {
		goto L130
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	v554 = v337
	goto L76
L140:
	;
	v546 = F_lappend(m, v337, v544)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v554 = v546
	goto L76
L142:
	;
	goto L75
L143:
	;
	goto L21
L144:
	;
	if int32(0) <= v665 {
		v101 = v665
		goto L19
	} else {
		goto L155
	}
L145:
	;
	v665 = base.I32_ctz(v651) | v652<<(uint(int32(5))%32)
	goto L144
L146:
	;
	v665 = int32(-2)
	goto L144
L147:
	;
	v616 = v101 + int32(1)
	v618 = base.I32_div_s(v616, int32(32))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v619 <= v618 {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v622 = v609 + int32(8)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622+v618<<(uint(int32(2))%32))))
	v629 = v626 & (int32(-1) << (uint(v616) % 32))
	if v629 != 0 {
		v651 = v629
		v652 = v618
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v631 = v618 + int32(1)
	if v631 == v619 {
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v634 = v631
	goto L151
L151:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v622+v634<<(uint(int32(2))%32))))
	if v641 != 0 {
		v651 = v641
		v652 = v634
		goto L145
	} else {
		goto L153
	}
L152:
	;
	goto L146
L153:
	;
	v643 = v634 + int32(1)
	if v643 != v619 {
		v634 = v643
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	goto L20
}
func F_generate_partitionwise_join_paths(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v8 - int32(1) {
	case 0, 2:
		goto L3
	default:
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+236)) = int32(0)
	return
L2:
	;
	return
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+232))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+240))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	if v20 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v25 == v23 {
		v45 = v23
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v45 != 0 {
		goto L2
	} else {
		goto L18
	}
L9:
	;
	goto L8
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v29 = v28
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if base.Ui32(int32(2)) <= base.Ui32(v33-int32(301)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = int32(1)
	goto L9
L13:
	;
	if v33 != int32(290) {
		v45 = v23
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v29 = v32 + int32(72)
	goto L11
L15:
	;
	goto L12
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v40 != 0 {
		v45 = v23
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+236))
	if int32(0) < v49 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_add_paths_to_append_rel(m, l0, l1, v99)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L47
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+252))
	v55 = v3
	v57 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_mark_dummy_rel(m, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L46
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52+v57<<(uint(int32(2))%32))))
	if v63 == int32(0) {
		v99 = v55
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v99 != 0 {
		goto L21
	} else {
		goto L45
	}
L27:
	;
	v101 = v57 + int32(1)
	if v101 != v49 {
		v55 = v99
		v57 = v101
		goto L25
	} else {
		goto L44
	}
L28:
	;
	F_generate_partitionwise_join_paths(m, l0, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	if v68 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_set_cheapest(m, v63)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	if v75 == v73 {
		v95 = v73
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v95 != 0 {
		v99 = v55
		goto L27
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v79 = v78
	goto L35
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if base.Ui32(int32(2)) <= base.Ui32(v83-int32(301)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v95 = int32(1)
	goto L33
L37:
	;
	if v83 != int32(290) {
		v95 = v73
		goto L33
	} else {
		goto L40
	}
L38:
	;
	v79 = v82 + int32(72)
	goto L35
L39:
	;
	goto L36
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+72))
	if v90 != 0 {
		v95 = v73
		goto L33
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v97 = F_lappend(m, v55, v63)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v99 = v97
	goto L27
L44:
	;
	goto L26
L45:
	;
	goto L24
L46:
	;
	return
L47:
	;
	F_list_free(m, v99)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	goto L2
}
func F_geqo_rand(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v30 float64
	_ = v30
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	v4 = v2 + int32(8)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
	v9 = v7 ^ v8
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = base.I64_rotl(v9, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = v9<<(uint(int64(16))%64) ^ base.I64_rotl(v7, int64(24)) ^ v9
	v30 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v7*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	return v30
}
func F_getTokenTypes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = F_lookup_ts_parser_cache(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		v196 = v3
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L49
	}
L5:
	;
	m.G0 = v14 + int32(32)
	return v196
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 == int32(0) {
		v196 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v28 = int32(0)
	v30 = F_OidFunctionCall1Coll(m, v25, v28, v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v32 <= int32(0) {
		v196 = v3
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v37 = v3
	v40 = v3
	v42 = v32
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v40<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v37 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v196 = v182
	goto L5
L13:
	;
	v192 = v40 + int32(1)
	if v192 < v187 {
		v37 = v182
		v40 = v192
		v42 = v187
		goto L11
	} else {
		goto L48
	}
L14:
	;
	if v30 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v54 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v59 = int32(0)
	goto L17
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57+v59<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v78 == int32(0) {
		v97 = v77
		v98 = v78
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L14
L19:
	;
	if v98-v97 == int32(0) {
		v182 = v37
		v187 = v42
		goto L13
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	if v77 != v78 {
		v97 = v77
		v98 = v78
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v82 = v51
	v83 = v74
	goto L23
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v97 = v86
		v98 = v87
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v97 = v86
	v98 = v87
	goto L20
L25:
	;
	v90 = int32(1)
	if v86 == v87 {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v103 = v59 + int32(1)
	if v54 != v103 {
		v59 = v103
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v119 == v118 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v122 = v118
	goto L31
L31:
	;
	v135 = v30 + v122*int32(12)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v140 == int32(0) {
		v159 = v139
		v160 = v140
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v169 = F_palloc0(m, int32(8))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	if v160-v159 != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	goto L33
L35:
	;
	if v139 != v140 {
		v159 = v139
		v160 = v140
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v144 = v51
	v145 = v136
	goto L37
L37:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v149 == int32(0) {
		v159 = v148
		v160 = v149
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v159 = v148
	v160 = v149
	goto L34
L39:
	;
	v152 = int32(1)
	if v148 == v149 {
		v144 = v144 + v152
		v145 = v145 + v152
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v163 = v122 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v30+v163*int32(12))))
	if v167 != 0 {
		v122 = v163
		goto L31
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L32
L44:
	;
	goto L3
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v174 = F_pstrdup(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v174
	v177 = F_lappend(m, v37, v169)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v182 = v177
	v187 = v179
	goto L13
L48:
	;
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(41185), v14)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(471337), int32(1243), int32(152722))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v240
	F_errmsg(m, int32(68905), v14+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(471337), int32(1278), int32(152722))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_actual_clauses(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v2 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v25
L2:
	;
	v8 = v2
	v9 = v2
	goto L7
L3:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v4 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v9<<(uint(int32(2))%32))))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = F_lappend(m, v8, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v25 = v16
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v21 = v9 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 < v22 {
		v8 = v16
		v9 = v21
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_get_actual_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 == v7 {
		v180 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 - int32(-64)
	return v180
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	if v20 == int32(0) {
		v180 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23+v24<<(uint(int32(2))%32))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+21)))
	if v29 == int32(112) {
		v180 = v7
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v32 <= int32(0) {
		v180 = v7
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = v7
	goto L6
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v43<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	if v52 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v180 = int32(0)
	goto L1
L8:
	;
	v168 = v43 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v168 < v169 {
		v43 = v168
		goto L6
	} else {
		goto L46
	}
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
	if v55 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+104)))
	if v56 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 != int32(1) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if l3 != v62 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v66 = F_match_index_to_operand(m, v64, int32(0), v51)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v66 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = F_get_op_opfamily_strategy(m, l2, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v102 = F_AllocSetContextCreateInternal(m, v97, int32(398184), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L28
	}
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v51)+64))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v51)+64))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v88 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = F_IndexAmTranslateStrategy(m, v74&int32(65535), v78, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	switch v81 - int32(1) {
	case 0:
		goto L19
	default:
		goto L8
	case 4:
		goto L18
	}
L22:
	;
	v89 = int32(-1)
	goto L24
L23:
	;
	v89 = int32(1)
	goto L24
L24:
	;
	v95 = v89
	goto L17
L25:
	;
	v94 = int32(1)
	goto L27
L26:
	;
	v94 = int32(-1)
	goto L27
L27:
	;
	v95 = v94
	goto L17
L28:
	;
	v104 = int32(4442576)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v102
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v110 = F_table_open(m, v108, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v114 = F_index_open(m, v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v117 = F_table_slot_create(m, v110, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_get_typlenbyval(m, v119, v13+int32(-2), v13+int32(-3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v126 = int32(1)
	v129 = int32(0)
	F_ScanKeyEntryInitialize(m, v15, int32(129), v126, v129, v129, v129, v129, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	if l4 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+62)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+61)))
	v138 = F_get_actual_variable_endpoint(m, v110, v114, v95, v15, v136, v137, v117, v105, l4)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L14
	} else {
		goto L37
	}
L35:
	;
	v140 = v126
	goto L36
L36:
	;
	v142 = base.B2i32(l5 == int32(0))
	v143 = v142 & v140
	if l5 == int32(0) {
		v154 = v143
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v140 = v138
	goto L36
L38:
	;
	F_ExecDropSingleTupleTableSlot(m, v117)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L14
	} else {
		goto L42
	}
L39:
	;
	if v140 == int32(0) {
		v154 = v143
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+62)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+61)))
	v152 = F_get_actual_variable_endpoint(m, v110, v114, int32(0)-v95, v15, v150, v151, v117, v105, l5)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v154 = v152
	goto L38
L42:
	;
	F_relation_close(m, v114, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	F_sequence_close(m, v110, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v105
	F_MemoryContextDelete(m, v102)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v180 = v154
	goto L1
L46:
	;
	goto L7
}
func F_get_attavgwidth(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1159]))
	if v5 != 0 {
		v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if int32(0) < v6 {
				v28 = v6
				return v28
			} else {
				v15 = F_SearchSysCache3(m, int32(65), l0, l1, int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					if v15 != 0 {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+12))
						F_ReleaseCatCache(m, v15)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							if int32(0) < v20 {
								v28 = v20
							} else {
								v28 = int32(0)
							}
							return v28
						}
					} else {
						v28 = int32(0)
						return v28
					}
				}
			}
		}
	} else {
		v15 = F_SearchSysCache3(m, int32(65), l0, l1, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+12))
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if int32(0) < v20 {
						v28 = v20
					} else {
						v28 = int32(0)
					}
					return v28
				}
			} else {
				v28 = int32(0)
				return v28
			}
		}
	}
}
func F_get_baserel_parampathinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 float64
	_ = v222
	var v223 int32
	_ = v223
	var v224 float64
	_ = v224
	var v232 float64
	_ = v232
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	v4 = int32(0)
	if l2 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v108 = F_bms_union(m, v107, l2)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = v4
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v25<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v39 = int32(0)
	v46 = base.B2i32(v38|l2 == v39)
	if v38 == v39 {
		v85 = v46
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v37
L9:
	;
	if v85 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	if l2 == int32(0) {
		v85 = v46
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v52 != v53 {
		v85 = int32(0)
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(1)
	if v52 <= v55 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v55
	goto L16
L15:
	;
	v58 = v52
	goto L16
L16:
	;
	v59 = int32(8)
	v64 = int32(0)
	goto L17
L17:
	;
	v72 = v64 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v38+v59+v72)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+(l2+v59))))
	v77 = base.B2i32(v74 == v76)
	if v76 != v74 {
		v85 = v77
		goto L10
	} else {
		goto L19
	}
L18:
	;
	v85 = v77
	goto L10
L19:
	;
	v80 = v64 + int32(1)
	if v80 != v58 {
		v64 = v80
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v92 = v25 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v92 < v93 {
		v25 = v92
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L8
L24:
	;
	goto L4
L25:
	;
	return int32(0)
L26:
	;
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v113 == v112 {
		v165 = v4
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v170 = F_generate_join_implied_equalities(m, l0, v108, l2, l1, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L25
	} else {
		goto L42
	}
L28:
	;
	v116 = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v117 <= v116 {
		v165 = v4
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v123 = v116
	v127 = v4
	goto L30
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v123<<(uint(int32(2))%32))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)+28))
	v139 = F_bms_is_subset(m, v138, v108)
	mBase = m.M
	if v139 == v137 {
		v150 = v137
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v165 = v153
	goto L27
L32:
	;
	if v150 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+28))
	v143 = F_bms_overlap(m, v136, v142)
	mBase = m.M
	if v143 == int32(0) {
		v150 = v137
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	v147 = F_bms_overlap(m, v136, v146)
	mBase = m.M
	v150 = v147 ^ int32(1)
	goto L33
L36:
	;
	v151 = F_lappend(m, v127, v135)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L25
	} else {
		goto L39
	}
L37:
	;
	v153 = v127
	goto L38
L38:
	;
	v155 = v123 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v155 < v156 {
		v123 = v155
		v127 = v153
		goto L30
	} else {
		goto L40
	}
L39:
	;
	v153 = v151
	goto L38
L40:
	;
	goto L31
L41:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v215 = F_list_concat_copy(m, v172, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L25
	} else {
		goto L50
	}
L42:
	;
	v172 = F_list_concat(m, v165, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	if v172 == int32(0) {
		v208 = v112
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v176 <= int32(0) {
		v208 = v112
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v183 = int32(0)
	v185 = v112
	goto L46
L46:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v183<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+56))
	v197 = F_bms_add_member(m, v185, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L25
	} else {
		goto L48
	}
L47:
	;
	v208 = v197
	goto L41
L48:
	;
	v200 = v183 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v200 < v201 {
		v183 = v200
		v185 = v197
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v217 = float64(1e+100)
	v218 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v220 = int32(0)
	v222 = F_clauselist_selectivity(m, l0, v215, v219, v220, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	v237 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_lt(v237, v236) != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v224 = base.F64_mul(v218, v222)
	if base.F64_gt(v224, float64(1e+100)) != 0 {
		v236 = v217
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v224)&int64(9223372036854775807)) {
		v236 = v217
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v232 = float64(1)
	if base.F64_le(v224, v232) != 0 {
		v236 = v232
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v236 = base.F64_nearest(v224)
	goto L51
L56:
	;
	v239 = v237
	goto L58
L57:
	;
	v239 = v236
	goto L58
L58:
	;
	v241 = F_palloc0(m, int32(24))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v241)+16)) = v172
	*(*float64)(unsafe.Add(mBase, uint32(v241)+8)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = int32(278)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v250 = F_lappend(m, v249, v241)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v250
	return v241
}
func F_get_cheapest_fractional_path(m *base.Module, l0 int32, l1 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.F64_le(l1, float64(0)) != 0 {
		v104 = v8
	} else {
		if base.F64_ge(l1, float64(1)) == int32(0) {
			v21 = l1
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)+32))
			if base.F64_gt(v15, float64(0)) == int32(0) {
				v21 = l1
			} else {
				v21 = base.F64_div(l1, v15)
			}
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if v23 == int32(0) {
			v104 = v8
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			if v26 <= int32(0) {
				v104 = v8
			} else {
				v31 = v8
				v33 = int32(0)
				for {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v33<<(uint(int32(2))%32))))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
					if v41 != 0 {
						v97 = v31
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if v40 == v42 {
							v97 = v31
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
							if v48 != v49 {
								if v48 < v49 {
									v54 = int32(-1)
								} else {
									v54 = int32(1)
								}
								v93 = v54
							} else {
								if base.F64_le(v21, float64(0))|base.F64_ge(v21, float64(1)) != 0 {
									v60 = int32(-1)
									v61 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
									v62 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
									if base.F64_lt(v61, v62) != 0 {
										v88 = v60
										v93 = v88
									} else {
										if base.F64_gt(v61, v62) != 0 {
											v93 = int32(1)
										} else {
											v66 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
											v67 = *(*float64)(unsafe.Add(mBase, uint32(v40)+48))
											if base.F64_lt(v66, v67) != 0 {
												v88 = v60
												v93 = v88
											} else {
												if base.F64_gt(v66, v67) != 0 {
													v88 = int32(1)
													v93 = v88
												} else {
													v93 = int32(0)
												}
											}
										}
									}
								} else {
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v31)+56))
									v74 = *(*float64)(unsafe.Add(mBase, uint32(v31)+48))
									v77 = base.F64_add(base.F64_mul(v21, base.F64_sub(v73, v74)), v74)
									v78 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
									v79 = *(*float64)(unsafe.Add(mBase, uint32(v40)+48))
									v82 = base.F64_add(base.F64_mul(v21, base.F64_sub(v78, v79)), v79)
									if base.F64_lt(v77, v82) != 0 {
										v88 = int32(-1)
									} else {
										v88 = base.F64_lt(v82, v77)
									}
									v93 = v88
								}
							}
							if v93 <= int32(0) {
								v96 = v31
							} else {
								v96 = v40
							}
							v97 = v96
						}
					}
					v99 = v33 + int32(1)
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if v99 < v100 {
						v31 = v97
						v33 = v99
						continue
					} else {
						break
					}
					break
				}
				v104 = v97
			}
		}
	}
	return v104
}
func F_get_cheapest_fractional_path_for_pathkeys(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	v4 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v212
L2:
	;
	v19 = v4
	v22 = v4
	goto L7
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v11 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v212 = v4
	goto L1
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v212 = v198
	goto L1
L9:
	;
	v204 = v22 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v204 < v205 {
		v19 = v198
		v22 = v204
		goto L7
	} else {
		goto L70
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	if v33 != v34 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if l1 == v81 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	if v78 <= int32(0) {
		v198 = v19
		goto L9
	} else {
		goto L31
	}
L14:
	;
	if v33 < v34 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v39 = int32(-1)
	goto L19
L18:
	;
	v39 = int32(1)
	goto L19
L19:
	;
	v78 = v39
	goto L13
L20:
	;
	v78 = v73
	goto L13
L21:
	;
	v45 = int32(-1)
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v28)+56))
	if base.F64_lt(v46, v47) != 0 {
		v73 = v45
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
	v62 = base.F64_add(base.F64_mul(l2, base.F64_sub(v58, v59)), v59)
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v28)+56))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v28)+48))
	v67 = base.F64_add(base.F64_mul(l2, base.F64_sub(v63, v64)), v64)
	if base.F64_lt(v62, v67) != 0 {
		v73 = int32(-1)
		goto L20
	} else {
		goto L30
	}
L24:
	;
	if base.F64_gt(v46, v47) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(1)
	goto L13
L26:
	;
	goto L27
L27:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v28)+48))
	if base.F64_lt(v51, v52) != 0 {
		v73 = v45
		goto L20
	} else {
		goto L28
	}
L28:
	;
	if base.F64_gt(v51, v52) != 0 {
		v73 = int32(1)
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v78 = int32(0)
	goto L13
L30:
	;
	v73 = base.F64_lt(v67, v62)
	goto L20
L31:
	;
	goto L12
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v133 != 0 {
		goto L50
	} else {
		goto L51
	}
L33:
	;
	v87 = int32(0)
	goto L34
L34:
	;
	v94 = int32(0)
	if l1 == v94 {
		v104 = v94
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v104 != 0 {
		v198 = v19
		goto L9
	} else {
		goto L49
	}
L36:
	;
	if v81 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v98 <= v87 {
		v104 = int32(0)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v104 = v100 + v87<<(uint(int32(2))%32)
	goto L36
L39:
	;
	if v104 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v87 < v105 {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v104 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v198 = v19
	goto L9
L45:
	;
	goto L35
L46:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v114 = v111 + v87<<(uint(int32(2))%32)
	if v114 == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v119 == v120 {
		v87 = v87 + int32(1)
		goto L34
	} else {
		goto L48
	}
L48:
	;
	v198 = v19
	goto L9
L49:
	;
	goto L32
L50:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v136 = v134
	goto L52
L51:
	;
	v136 = int32(0)
	goto L52
L52:
	;
	v137 = int32(0)
	if v136 == v137 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v191 != 0 {
		goto L67
	} else {
		goto L68
	}
L54:
	;
	v191 = int32(1)
	goto L53
L55:
	;
	goto L56
L56:
	;
	goto L57
L57:
	;
	v191 = v137
	goto L53
L67:
	;
	v192 = v28
	goto L69
L68:
	;
	v192 = v19
	goto L69
L69:
	;
	v198 = v192
	goto L9
L70:
	;
	goto L8
}
func F_get_cheapest_path_for_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	v6 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v223
L2:
	;
	v22 = v6
	v25 = v6
	goto L7
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v13 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v223 = v6
	goto L1
L6:
	;
	goto L5
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v223 = v207
	goto L1
L9:
	;
	v214 = v25 + int32(1)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v214 < v215 {
		v22 = v207
		v25 = v214
		goto L7
	} else {
		goto L80
	}
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+21)))
	if v33 != int32(1) {
		v207 = v22
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v22 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	if v40 != v41 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	if l1 == v86 {
		goto L42
	} else {
		goto L43
	}
L17:
	;
	if v83 <= int32(0) {
		v207 = v22
		goto L9
	} else {
		goto L41
	}
L18:
	;
	if v40 < v41 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if l3 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v46 = int32(-1)
	goto L23
L22:
	;
	v46 = int32(1)
	goto L23
L23:
	;
	v83 = v46
	goto L17
L24:
	;
	v83 = v77
	goto L17
L25:
	;
	v77 = int32(0)
	goto L24
L26:
	;
	v49 = int32(-1)
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v22)+48))
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v32)+48))
	if base.F64_lt(v50, v51) != 0 {
		v77 = v49
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v62 = int32(-1)
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v22)+56))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v32)+56))
	if base.F64_lt(v63, v64) != 0 {
		v77 = v62
		goto L24
	} else {
		goto L35
	}
L29:
	;
	if base.F64_gt(v50, v51) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = int32(1)
	goto L17
L31:
	;
	goto L32
L32:
	;
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v22)+56))
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v32)+56))
	if base.F64_lt(v55, v56) != 0 {
		v77 = v49
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if base.F64_gt(v55, v56) == int32(0) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v77 = int32(1)
	goto L24
L35:
	;
	if base.F64_gt(v63, v64) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v83 = int32(1)
	goto L17
L37:
	;
	goto L38
L38:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v22)+48))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v32)+48))
	if base.F64_lt(v68, v69) != 0 {
		v77 = v62
		goto L24
	} else {
		goto L39
	}
L39:
	;
	if base.F64_gt(v68, v69) != 0 {
		v77 = int32(1)
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L25
L41:
	;
	goto L16
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v142 != 0 {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v94 = int32(0)
	goto L44
L44:
	;
	v101 = int32(0)
	if l1 == v101 {
		v111 = v101
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v111 != 0 {
		v207 = v22
		goto L9
	} else {
		goto L59
	}
L46:
	;
	if v86 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v105 <= v94 {
		v111 = int32(0)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v111 = v107 + v94<<(uint(int32(2))%32)
	goto L46
L49:
	;
	if v111 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v94 < v112 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v111 == int32(0) {
		goto L42
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v207 = v22
	goto L9
L55:
	;
	goto L45
L56:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v121 = v118 + v94<<(uint(int32(2))%32)
	if v121 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v126 == v127 {
		v94 = v94 + int32(1)
		goto L44
	} else {
		goto L58
	}
L58:
	;
	v207 = v22
	goto L9
L59:
	;
	goto L42
L60:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	v145 = v143
	goto L62
L61:
	;
	v145 = int32(0)
	goto L62
L62:
	;
	v146 = int32(0)
	if v145 == v146 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v199 != 0 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v199 = int32(1)
	goto L63
L65:
	;
	goto L66
L66:
	;
	if l2 == int32(0) {
		v190 = v146
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v199 = v190
	goto L63
L68:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v156 < v155 {
		v190 = v146
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v158 = int32(1)
	if v155 <= v158 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v161 = v158
	goto L72
L71:
	;
	v161 = v155
	goto L72
L72:
	;
	v162 = int32(8)
	v167 = int32(0)
	goto L73
L73:
	;
	v174 = v167 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v145+v162+v174)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+(l2+v162))))
	v181 = v176 & (v178 ^ int32(-1))
	v183 = base.B2i32(v181 == int32(0))
	if v181 != 0 {
		v190 = v183
		goto L67
	} else {
		goto L75
	}
L74:
	;
	v190 = v183
	goto L67
L75:
	;
	v185 = v167 + int32(1)
	if v185 != v161 {
		v167 = v185
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v200 = v32
	goto L79
L78:
	;
	v200 = v22
	goto L79
L79:
	;
	v207 = v200
	goto L9
L80:
	;
	goto L8
}
func F_get_decomposed_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	v9 = l0 - int32(44032)
	if base.Ui32(v9) <= base.Ui32(int32(11171)) {
		v17 = base.I32_rem_u_s(v9&int32(65535), int32(28))
		if v17 != 0 {
			v18 = int32(3)
		} else {
			v18 = int32(2)
		}
		return v18
	} else {
		v20 = int32(1)
		v22 = l0 & int32(255)
		v23 = int32(8)
		v28 = int32(base.Ui32(l0<<(uint(v23)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
		v34 = int32(base.Ui32(int32(base.Ui32(l0)>>(uint(v23)%32))&int32(65280)) >> (uint(v23) % 32))
		v36 = int32(base.Ui32(l0) >> (uint(int32(24)) % 32))
		v37 = int32(8191)
		v48 = int32(13687)
		v49 = base.I32_rem_u_s(v22+(v28+(v34+v36*v37)*v37)*v37+int32(402620417), v48)
		v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49<<(uint(v20)%32))+uint32(_consts[1279]))))
		v55 = int32(257)
		v65 = base.I32_rem_u_s(((v36*v55+v34)*v55+v28)*v55+v22, v48)
		v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65<<(uint(v20)%32))+uint32(_consts[1279]))))
		v71 = v54 + v70
		if base.Ui32(int32(6842)) < base.Ui32(v71) {
			v256 = v20
		} else {
			v75 = v71 << (uint(int32(3)) % 32)
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[1280])))
			if l0 != v78 {
				v256 = v20
			} else {
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[1281]))))
				v82 = v80 & int32(31)
				if v82 == int32(0) {
					v256 = v20
				} else {
					if v80&int32(32) != 0 {
						v88 = l1
					} else {
						v88 = int32(1)
					}
					if v88 == int32(0) {
						v256 = v20
					} else {
						v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[1282]))))
						if v80&int32(64) != 0 {
							v94 = int32(4526096)
							*(*int32)(unsafe.Add(mBase, _consts[1283])) = v91
							v102 = int32(1)
							v103 = v94
						} else {
							v102 = v82
							v103 = v91<<(uint(int32(2))%32) + int32(2020928)
						}
						v104 = int32(0)
						v106 = v104
						v108 = v104
						for {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v103+v106<<(uint(int32(2))%32))))
							v123 = v116 - int32(44032)
							if base.Ui32(v123) <= base.Ui32(int32(11171)) {
								v131 = base.I32_rem_u_s(v123&int32(65535), int32(28))
								if v131 != 0 {
									v132 = int32(3)
								} else {
									v132 = int32(2)
								}
								v249 = v132
							} else {
								v133 = int32(1)
								v135 = v116 & int32(255)
								v136 = int32(8)
								v141 = int32(base.Ui32(v116<<(uint(v136)%32)&int32(16711680)) >> (uint(int32(16)) % 32))
								v147 = int32(base.Ui32(int32(base.Ui32(v116)>>(uint(v136)%32))&int32(65280)) >> (uint(v136) % 32))
								v149 = int32(base.Ui32(v116) >> (uint(int32(24)) % 32))
								v150 = int32(8191)
								v161 = int32(13687)
								v162 = base.I32_rem_u_s(v135+(v141+(v147+v149*v150)*v150)*v150+int32(402620417), v161)
								v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162<<(uint(v133)%32))+uint32(_consts[1279]))))
								v168 = int32(257)
								v178 = base.I32_rem_u_s(((v149*v168+v147)*v168+v141)*v168+v135, v161)
								v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v178<<(uint(v133)%32))+uint32(_consts[1279]))))
								v184 = v167 + v183
								if base.Ui32(int32(6842)) < base.Ui32(v184) {
									v237 = v133
								} else {
									v188 = v184 << (uint(int32(3)) % 32)
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1280])))
									if v116 != v191 {
										v237 = v133
									} else {
										v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1281]))))
										v195 = v193 & int32(31)
										if v195 == int32(0) {
											v237 = v133
										} else {
											if v193&int32(32) != 0 {
												v201 = l1
											} else {
												v201 = int32(1)
											}
											if v201 == int32(0) {
												v237 = v133
											} else {
												v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+uint32(_consts[1282]))))
												if v193&int32(64) != 0 {
													v207 = int32(4526096)
													*(*int32)(unsafe.Add(mBase, _consts[1283])) = v204
													v215 = int32(1)
													v216 = v207
												} else {
													v215 = v195
													v216 = v204<<(uint(int32(2))%32) + int32(2020928)
												}
												v217 = int32(0)
												v219 = v217
												v221 = v217
												for {
													v229 = *(*int32)(unsafe.Add(mBase, uint32(v216+v219<<(uint(int32(2))%32))))
													v230 = F_get_decomposed_size(m, v229, l1)
													mBase = m.M
													v231 = v230 + v221
													v233 = v219 + int32(1)
													if v233 != v215 {
														v219 = v233
														v221 = v231
														continue
													} else {
														break
													}
													break
												}
												v237 = v231
											}
										}
									}
								}
								v249 = v237
							}
							v250 = v249 + v108
							v252 = v106 + int32(1)
							if v252 != v102 {
								v106 = v252
								v108 = v250
								continue
							} else {
								break
							}
							break
						}
						v256 = v250
					}
				}
			}
		}
		return v256
	}
}
func F_get_element_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v5 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v15 = v13 + v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
				if v17 == int32(6179) {
					v21 = v16
				} else {
					v21 = int32(0)
				}
			} else {
				v21 = int32(0)
			}
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v21
			}
		}
	}
}
func F_get_formatted_log_time(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v2 = m.G0
	v4 = v2 - int32(32)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1203])))
	if v7 == int32(0) {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1204])))
		if v11 == int32(0) {
			F___gettimeofday(m, int32(4436160))
			mBase = m.M
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[1204])) = uint8(v17)
		} else {
		}
		v20 = *(*int64)(unsafe.Add(mBase, _consts[1205]))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v20
		v28 = *(*int32)(unsafe.Add(mBase, _consts[1206]))
		v29 = F_pg_localtime(m, v4+int32(24), v28)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = F_pg_strftime(m, int32(4436016), int32(128), int32(484445), v29)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1207]))
				v38 = base.I32_div_s(v36, int32(1000))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v38
				v43 = F_pg_sprintf(m, v4+int32(8), int32(443944), v4)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					*(*int32)(unsafe.Add(mBase, _consts[1208])) = v46
					m.G0 = v4 + int32(32)
					return int32(4436016)
				}
			}
		}
	} else {
		m.G0 = v4 + int32(32)
		return int32(4436016)
	}
}
func F_get_formatted_start_time(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
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
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int64)(unsafe.Add(mBase, _consts[1209]))
	*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1210])))
	if v10 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[1206]))
		v20 = F_pg_localtime(m, v4+int32(8), v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_pg_strftime(m, int32(4436176), int32(128), int32(484424), v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v4 + int32(16)
				return int32(4436176)
			}
		}
	} else {
		m.G0 = v4 + int32(16)
		return int32(4436176)
	}
}
func F_get_nullingrels_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v12 = l0
	v13 = l1
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 != int32(64) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L1
L5:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v108 != 0 {
		v12 = v108
		v13 = v106
		goto L3
	} else {
		goto L39
	}
L6:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_get_nullingrels_recurse(m, v103, v102, l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L17
	} else {
		goto L38
	}
L7:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95+v96<<(uint(int32(2))%32)))) = v13
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L35
	}
L9:
	;
	switch v17 - int32(63) {
	case 0:
		goto L7
	default:
		goto L8
	case 2:
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	switch v45 {
	case 0:
		v61 = v13
		goto L21
	case 1, 4, 5:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	default:
		goto L20
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v29 = int32(0)
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v29<<(uint(int32(2))%32))))
	F_get_nullingrels_recurse(m, v38, v13, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L1
L17:
	;
	return
L18:
	;
	v42 = v29 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v42 < v43 {
		v29 = v42
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L32
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_get_nullingrels_recurse(m, v62, v61, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L31
	}
L22:
	;
	v56 = F_bms_copy(m, v13)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L29
	}
L23:
	;
	v51 = F_bms_copy(m, v13)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L27
	}
L24:
	;
	v46 = F_bms_copy(m, v13)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v49 = F_bms_add_member(m, v46, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v101 = v49
	v102 = v13
	goto L6
L27:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v54 = F_bms_add_member(m, v51, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v101 = v54
	v102 = v54
	goto L6
L29:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v59 = F_bms_add_member(m, v56, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v61 = v59
	goto L21
L31:
	;
	v106 = v13
	goto L5
L32:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
	F_errmsg_internal(m, int32(462600), v8+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(475999), int32(4444), int32(343085))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v85
	F_errmsg_internal(m, int32(463405), v8)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(475999), int32(4450), int32(343085))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v106 = v101
	goto L5
L39:
	;
	goto L4
}
func F_get_oprrest(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+104))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_ordering_op_for_equality_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v2 = int32(0)
	v11 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v2, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_ReleaseCatCacheList(m, v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L26
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v15 <= int32(0) {
		v77 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = int32(0)
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(48)+v21<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v34 <= int32(2741) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v77 = int32(0)
	goto L1
L7:
	;
	v72 = v21 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v72 < v73 {
		v21 = v72
		goto L5
	} else {
		goto L25
	}
L8:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+16)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v60 = F_IndexAmTranslateStrategy(m, v58, v56, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L21
	}
L9:
	;
	v48 = F_GetIndexAmRoutineByAmId(m, v34, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L18
	}
L10:
	;
	switch v34 - int32(403) {
	case 0:
		v56 = v34
		goto L8
	case 1:
		goto L9
	case 2:
		goto L7
	default:
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v34 == int32(2742) {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	if v34 != int32(783) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	if v34 == int32(3580) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if v34 == int32(4000) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+10)))
	F_pfree(m, v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v50 != int32(1) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v56 = v55
	goto L8
L21:
	;
	if v60 != int32(3) {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v67 = F_get_opfamily_member_for_cmptype(m, v64, v65, v65, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v67 != 0 {
		v77 = v67
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	goto L6
L26:
	;
	return v77
}
func F_get_relname_relid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_GetSysCacheOid(m, int32(56), l0, l1, v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_get_required_extension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v18 = F_GetSysCacheOid(m, int32(27), l0, v7, v7, v7)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L37
	}
L3:
	;
	return int32(0)
L4:
	;
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v122 = v18
	goto L7
L7:
	;
	m.G0 = v12 + int32(48)
	return v122
L8:
	;
	F_check_valid_extension_name(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if l4 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v93 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L28
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v30 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(0)
	if v33 < v30 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = v30
	goto L15
L14:
	;
	v37 = v33
	goto L15
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v46 = v33
	goto L16
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38+v46<<(uint(int32(2))%32))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == int32(0) {
		v74 = v54
		v75 = v55
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L10
L18:
	;
	if v75-v74 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	if v54 != v55 {
		v74 = v54
		v75 = v55
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v59 = v51
	v60 = l0
	goto L22
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v64 == int32(0) {
		v74 = v63
		v75 = v64
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v74 = v63
	v75 = v64
	goto L19
L24:
	;
	v67 = int32(1)
	if v63 == v64 {
		v59 = v59 + v67
		v60 = v60 + v67
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v80 = v46 + int32(1)
	if v80 != v37 {
		v46 = v80
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	if v93 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(667927), v12)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v108 = F_list_copy(m, l4)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L34
	}
L32:
	;
	F_errfinish(m, int32(473200), int32(2059), int32(258723))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v110 = F_lappend(m, v108, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_CreateExtensionInternal(m, v12+int32(36), l0, l2, int32(0), int32(1), v110, l5)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v122 = v114
	goto L7
L37:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg(m, int32(680891), v12+int32(16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(473200), int32(2054), int32(258723))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg(m, int32(432649), v12+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	if l5 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_errhint(m, int32(571663), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errfinish(m, int32(473200), int32(2084), int32(258723))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_sortgroupref_clause_noerr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v38
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 <= int32(0) {
		v38 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = int32(0)
	goto L1
L5:
	;
	v9 = int32(0)
	if v9 < v6 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v12 = v6
	goto L8
L7:
	;
	v12 = v9
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = int32(0)
	goto L9
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13+v16<<(uint(int32(2))%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == l0 {
		v38 = v23
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L4
L11:
	;
	v27 = v16 + int32(1)
	if v27 != v12 {
		v16 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_get_steps_using_prefix_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v22 = int32(2)
	v23 = (l7 - v20) >> (uint(v22) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20+v26<<(uint(v22)%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v25 < v33 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v182
L4:
	;
	if v26 <= v23 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v110 = int32(0)
	if v26 <= v23 {
		v182 = v110
		goto L3
	} else {
		goto L27
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L9
L9:
	;
	v45 = v23
	goto L11
L10:
	;
	v76 = int32(0)
	v77 = v23
	goto L15
L11:
	;
	v55 = v20 + v45<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v25 < v57 {
		v64 = v55
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v64 = int32(0)
	goto L10
L13:
	;
	v60 = v45 + int32(1)
	if v60 != v26 {
		v45 = v60
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v77<<(uint(int32(2))%32))))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != v25 {
		v182 = v76
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v182 = v100
	goto L3
L17:
	;
	v88 = F_list_copy(m, l8)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v91 = F_lappend(m, v88, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v93 = F_list_copy(m, l9)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v96 = F_lappend_oid(m, v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v98 = F_get_steps_using_prefix_recurse(m, l0, l1, l2, l3, l4, l5, l6, v64, v91, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v100 = F_list_concat(m, v76, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_list_free(m, v91)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_list_free(m, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v107 = v77 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v107 < v108 {
		v76 = v100
		v77 = v107
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	if l2 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v113 = int32(0)
	goto L30
L29:
	;
	v113 = l1
	goto L30
L30:
	;
	v124 = v110
	v125 = v23
	goto L31
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+v125<<(uint(int32(2))%32))))
	v134 = F_list_copy(m, l8)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v182 = v166
	goto L3
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v137 = F_lappend(m, v134, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v139 = F_lappend(m, v137, l3)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v141 = F_list_copy(m, l9)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v144 = F_lappend_oid(m, v141, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v146 = F_lappend_oid(m, v144, l4)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v149 = F_palloc0(m, int32(24))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(377)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v153 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v139
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+8)) = uint16(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v153
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v163 = F_lappend(m, v162, v149)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v163
	v166 = F_lappend(m, v124, v149)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v169 = v125 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v169 < v170 {
		v124 = v166
		v125 = v169
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L32
}
func F_get_typavgwidth(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	v6 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10+v11)+76)))
			F_ReleaseCatCache(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if int32(0) < v13 {
					v81 = v13
					return v81
				} else {
					v20 = int32(-1)
					if l1 < int32(0) {
						v60 = v20
						v62 = v60
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
							switch l0 - int32(1560) {
							case 0, 2:
								v56 = int32(8)
								v57 = base.I32_div_s(l1+int32(7), v56)
								v60 = v57 + v56
								v62 = v60
							case 1:
								v60 = v20
								v62 = v60
							default:
								if l0 != int32(1700) {
									v60 = v20
									v62 = v60
								} else {
									if int32(4) <= l1 {
										v53 = int32(base.Ui32(int32(base.Ui32(l1-int32(4))>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(65534) + int32(8)
									} else {
										v53 = int32(-1)
									}
									v62 = v53
								}
							}
						} else {
							v29 = F_GetDatabaseEncoding(m)
							mBase = m.M
							v30 = F_pg_encoding_max_length(m, v29)
							mBase = m.M
							v31 = int32(4)
							v62 = v30*(l1-v31) + v31
						}
					}
					if v62 <= int32(0) {
						return int32(32)
					} else {
						if l0 == int32(1042) {
							v81 = v62
							return v81
						} else {
							if base.Ui32(v62) < base.Ui32(int32(33)) {
								v81 = v62
								return v81
							} else {
								if base.Ui32(int32(999)) < base.Ui32(v62) {
									return int32(516)
								} else {
									v75 = int32(32)
									v81 = int32(base.Ui32(v62-v75)>>(uint(int32(1))%32)) + v75
									return v81
								}
							}
						}
					}
				}
			}
		} else {
			v20 = int32(-1)
			if l1 < int32(0) {
				v60 = v20
				v62 = v60
			} else {
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(1042)) {
					switch l0 - int32(1560) {
					case 0, 2:
						v56 = int32(8)
						v57 = base.I32_div_s(l1+int32(7), v56)
						v60 = v57 + v56
						v62 = v60
					case 1:
						v60 = v20
						v62 = v60
					default:
						if l0 != int32(1700) {
							v60 = v20
							v62 = v60
						} else {
							if int32(4) <= l1 {
								v53 = int32(base.Ui32(int32(base.Ui32(l1-int32(4))>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(65534) + int32(8)
							} else {
								v53 = int32(-1)
							}
							v62 = v53
						}
					}
				} else {
					v29 = F_GetDatabaseEncoding(m)
					mBase = m.M
					v30 = F_pg_encoding_max_length(m, v29)
					mBase = m.M
					v31 = int32(4)
					v62 = v30*(l1-v31) + v31
				}
			}
			if v62 <= int32(0) {
				return int32(32)
			} else {
				if l0 == int32(1042) {
					v81 = v62
					return v81
				} else {
					if base.Ui32(v62) < base.Ui32(int32(33)) {
						v81 = v62
						return v81
					} else {
						if base.Ui32(int32(999)) < base.Ui32(v62) {
							return int32(516)
						} else {
							v75 = int32(32)
							v81 = int32(base.Ui32(v62-v75)>>(uint(int32(1))%32)) + v75
							return v81
						}
					}
				}
			}
		}
	}
}
func F_get_useful_group_keys_orderings(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v17 = F_palloc0(m, int32(12))
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(276)
	v26 = F_lappend(m, int32(0), v17)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[487])))
	if v29 != int32(1) {
		v309 = v26
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v309
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v32 != 0 {
		v309 = v26
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v33 == int32(0) {
		v309 = v26
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v33 == v36 {
		v309 = v26
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v42 = int32(0)
	goto L10
L9:
	;
	if v14 == int32(0) {
		v309 = v26
		goto L4
	} else {
		goto L26
	}
L10:
	;
	if v42 < v38 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v58 == int32(0) {
		v309 = v26
		goto L4
	} else {
		goto L25
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v58 = v54 + v42<<(uint(int32(2))%32)
	goto L14
L13:
	;
	v58 = int32(0)
	goto L14
L14:
	;
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v58 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v42 < v59 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v58 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v309 = v26
	goto L4
L21:
	;
	goto L11
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v66 = v63 + v42<<(uint(int32(2))%32)
	if v66 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v71 == v72 {
		v42 = v42 + int32(1)
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L9
L25:
	;
	goto L9
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v82 = F_list_copy_head(m, v14, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v85 <= v84 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v229 = F_list_concat_unique_ptr(m, v217, v14)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L74
	}
L29:
	;
	v88 = int32(0)
	v217 = v84
	v223 = v88
	v228 = v88
	goto L28
L30:
	;
	goto L31
L31:
	;
	v90 = int32(0)
	if v90 < v81 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = v81
	goto L34
L33:
	;
	v93 = v90
	goto L34
L34:
	;
	v94 = int32(0)
	v97 = v84
	v98 = v94
	v103 = v94
	goto L35
L35:
	;
	if v98 == v93 {
		v206 = v97
		v209 = v103
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v206 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L37:
	;
	goto L36
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v98<<(uint(int32(2))%32))))
	v114 = int32(0)
	if v82 == v114 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v152 == int32(0) {
		v206 = v97
		v209 = v103
		goto L37
	} else {
		goto L52
	}
L40:
	;
	v152 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v120 <= int32(0) {
		v145 = v114
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = v145
	goto L39
L44:
	;
	v123 = int32(0)
	if v123 < v120 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v120
	goto L47
L46:
	;
	v126 = v123
	goto L47
L47:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v129 = int32(0)
	goto L48
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127+v129<<(uint(int32(2))%32))))
	v138 = base.B2i32(v137 == v113)
	if v137 == v113 {
		v145 = v138
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v145 = v138
	goto L43
L50:
	;
	v140 = v129 + int32(1)
	if v140 != v126 {
		v129 = v140
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+44))
	if v156 == int32(0) {
		v206 = v97
		v209 = v103
		goto L37
	} else {
		goto L53
	}
L53:
	;
	if v15 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v194 == int32(0) {
		v206 = v97
		v209 = v103
		goto L37
	} else {
		goto L67
	}
L55:
	;
	goto L54
L56:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v162 <= int32(0) {
		v194 = int32(0)
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v194 = int32(0)
	goto L55
L59:
	;
	v165 = int32(0)
	if v165 < v162 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v168 = v162
	goto L62
L61:
	;
	v168 = v165
	goto L62
L62:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v172 = int32(0)
	goto L63
L63:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v169+v172<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v180 == v156 {
		v194 = v179
		goto L55
	} else {
		goto L65
	}
L64:
	;
	goto L58
L65:
	;
	v183 = v172 + int32(1)
	if v183 != v168 {
		v172 = v183
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v198 = F_lappend(m, v97, v113)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v200 = F_lappend(m, v103, v194)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v203 = v98 + int32(1)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v203 < v204 {
		v97 = v198
		v98 = v203
		v103 = v200
		goto L35
	} else {
		goto L70
	}
L70:
	;
	v206 = v198
	v209 = v200
	goto L37
L71:
	;
	v213 = int32(0)
	v217 = v213
	v223 = v209
	v228 = v213
	goto L28
L72:
	;
	goto L73
L73:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v217 = v206
	v223 = v209
	v228 = v215
	goto L28
L74:
	;
	v231 = F_list_concat_unique_ptr(m, v223, v15)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_list_free(m, v82)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v228 <= int32(0) {
		v309 = v26
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _consts[488])))
	if v238 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v228 != v241 {
		v309 = v26
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v229 == v243 {
		v309 = v26
		goto L4
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v246 = int32(0)
	goto L85
L83:
	;
	v292 = F_palloc0(m, int32(12))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L98
	}
L84:
	;
	if v268 == int32(0) {
		v309 = v26
		goto L4
	} else {
		goto L97
	}
L85:
	;
	v258 = int32(0)
	if v229 == v258 {
		v268 = v258
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v268|v276 != 0 {
		goto L83
	} else {
		goto L96
	}
L87:
	;
	if v243 == int32(0) {
		goto L84
	} else {
		goto L90
	}
L88:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v262 <= v246 {
		v268 = int32(0)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v268 = v264 + v246<<(uint(int32(2))%32)
	goto L87
L90:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v271 <= v246 {
		goto L84
	} else {
		goto L91
	}
L91:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v276 = v273 + v246<<(uint(int32(2))%32)
	if v268 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	goto L86
L93:
	;
	if v276 == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v283 == v284 {
		v246 = v246 + int32(1)
		goto L85
	} else {
		goto L95
	}
L95:
	;
	goto L83
L96:
	;
	v309 = v26
	goto L4
L97:
	;
	goto L83
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = int32(276)
	v298 = F_lappend(m, v26, v292)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v309 = v298
	goto L4
}
func F_get_useful_pathkeys_for_distinct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	v4 = int32(0)
	v9 = F_lappend(m, v4, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[506])))
	if v14 != int32(1) {
		v213 = v9
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v213
L4:
	;
	if l2 == int32(0) {
		v213 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v19 <= int32(0) {
		v131 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v131 == int32(0) {
		v213 = v9
		goto L3
	} else {
		goto L43
	}
L7:
	;
	v25 = v4
	v26 = v4
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v25<<(uint(int32(2))%32))))
	v34 = int32(0)
	if l1 == v34 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v131 = v121
	goto L6
L10:
	;
	if v72 == int32(0) {
		v131 = v26
		goto L6
	} else {
		goto L23
	}
L11:
	;
	v72 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v40 <= int32(0) {
		v65 = v34
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v72 = v65
	goto L10
L15:
	;
	v43 = int32(0)
	if v43 < v40 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v46 = v40
	goto L18
L17:
	;
	v46 = v43
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = int32(0)
	goto L19
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47+v49<<(uint(int32(2))%32))))
	v58 = base.B2i32(v57 == v33)
	if v57 == v33 {
		v65 = v58
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v65 = v58
	goto L14
L21:
	;
	v60 = v49 + int32(1)
	if v60 != v46 {
		v49 = v60
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+40)))
	if v76 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v80 = int32(0)
	if v79 == v80 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v121 = F_lappend(m, v26, v33)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L41
	}
L27:
	;
	if v118 == int32(0) {
		v131 = v26
		goto L6
	} else {
		goto L40
	}
L28:
	;
	v118 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v86 <= int32(0) {
		v111 = v80
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = v111
	goto L27
L32:
	;
	v89 = int32(0)
	if v89 < v86 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = v86
	goto L35
L34:
	;
	v92 = v89
	goto L35
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v95 = int32(0)
	goto L36
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93+v95<<(uint(int32(2))%32))))
	v104 = base.B2i32(v103 == v33)
	if v103 == v33 {
		v111 = v104
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v111 = v104
	goto L31
L38:
	;
	v106 = v95 + int32(1)
	if v106 != v92 {
		v95 = v106
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L26
L41:
	;
	v124 = v25 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v124 < v125 {
		v25 = v124
		v26 = v121
		goto L8
	} else {
		goto L42
	}
L42:
	;
	goto L9
L43:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v139 = v137
	goto L46
L45:
	;
	v139 = int32(0)
	goto L46
L46:
	;
	if v136 < v139 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, _consts[488])))
	if v142 != int32(1) {
		v213 = v9
		goto L3
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v145 = F_list_concat_unique_ptr(m, v131, l1)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	if l1 == v145 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v203 == int32(0) {
		v213 = v9
		goto L3
	} else {
		goto L74
	}
L53:
	;
	v203 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v154 = int32(0)
	goto L58
L56:
	;
	if v193 != 0 {
		goto L71
	} else {
		goto L72
	}
L57:
	;
	v187 = int32(0)
	v193 = base.B2i32(v167 == v187)
	v195 = base.B2i32(v176 != v187) << (uint(int32(1)) % 32)
	goto L56
L58:
	;
	v157 = int32(0)
	if l1 == v157 {
		v167 = v157
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v203 = int32(3)
	goto L52
L60:
	;
	if v145 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v161 <= v154 {
		v167 = int32(0)
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v167 = v163 + v154<<(uint(int32(2))%32)
	goto L60
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	v176 = v173 + v154<<(uint(int32(2))%32)
	if v167 == int32(0) {
		goto L57
	} else {
		goto L68
	}
L64:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v154 < v168 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v170 = int32(0)
	v193 = base.B2i32(v167 == v170)
	v195 = v170
	goto L56
L67:
	;
	goto L66
L68:
	;
	if v176 == int32(0) {
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v183 == v184 {
		v154 = v154 + int32(1)
		goto L58
	} else {
		goto L70
	}
L70:
	;
	goto L59
L71:
	;
	v197 = v195
	goto L73
L72:
	;
	v197 = int32(1)
	goto L73
L73:
	;
	v203 = v197
	goto L52
L74:
	;
	v206 = F_lappend(m, v9, v145)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v213 = v206
	goto L3
}
func F_get_windowfunc_expr_helper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(448)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v18 == v6 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L19
	}
L3:
	;
	v32 = v6
	v37 = v6
	goto L10
L4:
	;
	v82 = v6
	v87 = v6
	goto L1
L5:
	;
	goto L4
L6:
	;
	goto L7
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if int32(101) <= v21 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if int32(0) < v24 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v40 = v32 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 == int32(16) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v82 = v58
	v87 = v50
	goto L1
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v48 = F_lappend(m, v37, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v50 = v37
	goto L14
L14:
	;
	v54 = F_exprType(m, v43)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v50 = v48
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(48)+v40))) = v54
	v58 = v32 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v58 < v59 {
		v32 = v58
		v37 = v50
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(112865), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(470833), int32(11049), int32(207574))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v97 = l2
	goto L25
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v92 = int32(0)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	v95 = F_generate_function_name(m, v89, v82, v87, v15+int32(48), v92, v92, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v97
	F_appendStringInfo(m, v17, int32(643629), v15+int32(32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L27
	}
L26:
	;
	v97 = v95
	goto L25
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v104 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if l3 != 0 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	F_appendStringInfoChar(m, v17, int32(42))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l4 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	F_get_rule_expr(m, v112, l1, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_get_rule_expr(m, v110, l1, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L15
	} else {
		goto L39
	}
L36:
	;
	F_appendStringInfoString(m, v17, int32(703732))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	F_get_rule_expr(m, v121, l1, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	goto L28
L39:
	;
	goto L28
L40:
	;
	F_appendStringInfoString(m, v17, l3)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v131 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_appendStringInfoString(m, v17, int32(702811))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L15
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_appendStringInfoString(m, v17, int32(702230))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L49
	}
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_get_rule_expr(m, v135, l1, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v142 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	m.G0 = v15 + int32(448)
	return
L51:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v233)+72))
	v275 = F_quote_identifier(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L15
	} else {
		goto L87
	}
L52:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_get_rule_windowspec(m, v164, v271, l1)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L15
	} else {
		goto L86
	}
L53:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v143 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v205 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L67
	}
L57:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v154 = int32(0)
	goto L58
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v147+v154<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	if v146 != v165 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v170 == int32(0) {
		goto L52
	} else {
		goto L64
	}
L60:
	;
	v168 = v154 + int32(1)
	if v143 != v168 {
		v154 = v168
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L56
L64:
	;
	v173 = F_quote_identifier(m, v170)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	F_appendStringInfoString(m, v17, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	goto L50
L67:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v193
	F_errmsg_internal(m, int32(46434), v15+int32(16))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(470833), int32(11113), int32(207574))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L15
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L83
	}
L71:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v208 <= int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v211 = int32(0)
	if v211 < v208 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v215 = v208
	goto L75
L74:
	;
	v215 = v211
	goto L75
L75:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v222 = v211
	goto L76
L76:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v216+v222<<(uint(int32(2))%32))))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+40))
	if v233 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L70
L78:
	;
	v243 = v222 + int32(1)
	if v243 != v215 {
		v222 = v243
		goto L76
	} else {
		goto L82
	}
L79:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v236 != int32(366) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+76))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v239 == v240 {
		goto L51
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	goto L77
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v261
	F_errmsg_internal(m, int32(46434), v15)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(470833), int32(11138), int32(207574))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	goto L50
L87:
	;
	F_appendStringInfoString(m, v17, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L15
	} else {
		goto L88
	}
L88:
	;
	goto L50
}
func F_get_xmltable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v155 int32
	_ = v155
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v361 int32
	_ = v361
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v24, int32(643975))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, v24, int32(644218))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_appendStringInfoChar(m, v24, int32(40))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L42
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 == int32(0) {
		v40 = v4
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		v40 = v4
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v40 = v39
	goto L7
L10:
	;
	F_appendStringInfoString(m, v24, int32(703945))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L41
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v43 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v40 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v48 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v73 = int32(1)
	goto L24
L16:
	;
	F_get_rule_expr(m, v51, l1, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_appendStringInfoString(m, v24, int32(701996))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v56 = F_quote_identifier(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v56
	F_appendStringInfo(m, v24, int32(187739), v20+int32(-16))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	F_get_rule_expr(m, v51, l1, l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v89 = int32(0)
	if v33 == v89 {
		v98 = v89
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v99 <= v73 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v92 <= v73 {
		v98 = v89
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v98 = v94 + v73<<(uint(int32(2))%32)
	goto L26
L29:
	;
	if v98 == int32(0) {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v106 = v103 + v73<<(uint(int32(2))%32)
	if v106 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	F_appendStringInfoString(m, v24, int32(703976))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v109 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_get_rule_expr(m, v110, l1, l2)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoString(m, v24, int32(701996))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v117 = F_quote_identifier(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v117
	F_appendStringInfo(m, v24, int32(187739), v20+int32(-32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v73 = v73 + int32(1)
	goto L24
L39:
	;
	F_get_rule_expr(m, v110, l1, l2)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v73 = v73 + int32(1)
	goto L24
L41:
	;
	goto L5
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_rule_expr(m, v178, l1, l2)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_appendStringInfoString(m, v24, int32(644923))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_rule_expr(m, v184, l1, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v190 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L98
	}
L48:
	;
	F_appendStringInfoString(m, v24, int32(702115))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v205 = int32(0)
	goto L50
L50:
	;
	v221 = int32(0)
	if v200 == v221 {
		v231 = v221
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v232 = int32(0)
	if v199 == v232 {
		v243 = v232
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v225 <= v205 {
		v231 = int32(0)
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v231 = v227 + v205<<(uint(int32(2))%32)
	goto L52
L55:
	;
	if v198 == int32(0) {
		v252 = v232
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v237 <= v205 {
		v243 = int32(0)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v243 = v239 + v205<<(uint(int32(2))%32)
	goto L55
L58:
	;
	v253 = int32(0)
	if v197 == v253 {
		v264 = v253
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v246 <= v205 {
		v252 = v232
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v252 = v248 + v205<<(uint(int32(2))%32)
	goto L58
L61:
	;
	if v196 == int32(0) {
		v273 = v253
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v258 <= v205 {
		v264 = int32(0)
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v264 = v260 + v205<<(uint(int32(2))%32)
	goto L61
L64:
	;
	if v231 == int32(0) {
		goto L47
	} else {
		goto L67
	}
L65:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v267 <= v205 {
		v273 = v253
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v273 = v269 + v205<<(uint(int32(2))%32)
	goto L64
L67:
	;
	if v243 == int32(0) {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	if v252 == int32(0) {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	if v264 == int32(0) {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	if v273 == int32(0) {
		goto L47
	} else {
		goto L71
	}
L71:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v292 = F_bms_is_member(m, v205, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if int32(0) < v205 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_appendStringInfoString(m, v24, int32(703976))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v300 = v205 + int32(1)
	v301 = F_quote_identifier(m, v290)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	if v205 != v284 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v304 = F_format_type_with_typemod(m, v288, v287)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(484702)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v301
	F_appendStringInfo(m, v24, int32(170509), v22)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L97
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v301
	F_appendStringInfo(m, v24, int32(170509), v20+int32(-48))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v285 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_appendStringInfoString(m, v24, int32(644158))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v286 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	F_get_rule_expr(m, v285, l1, l2)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L85
L89:
	;
	F_appendStringInfoString(m, v24, int32(644906))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v292 == int32(0) {
		v205 = v300
		goto L50
	} else {
		goto L95
	}
L92:
	;
	F_get_rule_expr(m, v286, l1, l2)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_appendStringInfoChar(m, v24, int32(41))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	F_appendStringInfoString(m, v24, int32(508928))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v205 = v300
	goto L50
L97:
	;
	v205 = v300
	goto L50
L98:
	;
	m.G0 = v22 - int32(-64)
	return
}
func F_getcwd(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = m.G0
	v8 = int32(4096)
	if l0 != 0 {
		v11 = int32(16)
	} else {
		v11 = v8
	}
	v12 = v7 - v11
	m.G0 = v12
	if l0 == int32(0) {
		v20 = v12
		v21 = v8
		v22 = int32(0)
		v23 = m.Env.X__syscall_getcwd(m, v20, v21)
		mBase = m.M
		if base.Ui32(int32(-4095)) <= base.Ui32(v23) {
			*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0) - v23
			v31 = int32(-1)
		} else {
			v31 = v23
		}
		if v31 < int32(0) {
			v52 = v22
		} else {
			if v31 != 0 {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v34 == int32(47) {
					if v20 != v12 {
						v52 = v20
					} else {
						v43 = F_strlen(m, v20)
						mBase = m.M
						v45 = v43 + int32(1)
						v46 = F_emscripten_builtin_malloc(m, v45)
						mBase = m.M
						if v46 == int32(0) {
							v51 = int32(0)
						} else {
							v50 = F___memcpy(m, v46, v20, v45)
							mBase = m.M
							v51 = v50
						}
						v52 = v51
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(44)
					v52 = v22
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(44)
				v52 = v22
			}
		}
	} else {
		if l1 != 0 {
			v20 = l0
			v21 = l1
			v22 = int32(0)
			v23 = m.Env.X__syscall_getcwd(m, v20, v21)
			mBase = m.M
			if base.Ui32(int32(-4095)) <= base.Ui32(v23) {
				*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0) - v23
				v31 = int32(-1)
			} else {
				v31 = v23
			}
			if v31 < int32(0) {
				v52 = v22
			} else {
				if v31 != 0 {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
					if v34 == int32(47) {
						if v20 != v12 {
							v52 = v20
						} else {
							v43 = F_strlen(m, v20)
							mBase = m.M
							v45 = v43 + int32(1)
							v46 = F_emscripten_builtin_malloc(m, v45)
							mBase = m.M
							if v46 == int32(0) {
								v51 = int32(0)
							} else {
								v50 = F___memcpy(m, v46, v20, v45)
								mBase = m.M
								v51 = v50
							}
							v52 = v51
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(44)
						v52 = v22
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(44)
					v52 = v22
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(28)
			v52 = int32(0)
		}
	}
	m.G0 = v7
	return v52
}
func F_getdatabaseencoding(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[358]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_geterrcode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1184]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1184])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(432766), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(474769), int32(1588), int32(393940))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_consts[1195])))
		return v28
	}
}
func F_geterrposition(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1184]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1184])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(432766), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(474769), int32(1605), int32(236837))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_consts[1202])))
		return v28
	}
}
func F_gettoken_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(160)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = int32(1)
	v30 = v21
	v32 = v7
	v33 = v7
	v36 = v7
	v38 = v7
	goto L1
L1:
	;
	switch v29 - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L27
	case 2:
		goto L28
	case 3:
		goto L26
	case 4:
		v74 = v30
		goto L30
	case 5:
		goto L21
	case 6:
		goto L20
	case 7:
		goto L31
	default:
		goto L19
	}
L3:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v814 = F_pg_mblen_cstr(m, v813)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L40
	} else {
		goto L329
	}
L4:
	;
	v804 = int32(3)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v802
	goto L3
L5:
	;
	v802 = int32(2)
	goto L4
L6:
	;
	m.G0 = v19 + int32(160)
	return v783
L7:
	;
	v767 = int32(1)
	if base.Ui32(v83-int32(9)) < base.Ui32(int32(5)) {
		v804 = v767
		v805 = v30
		v807 = v32
		v808 = v33
		v810 = v36
		v812 = v38
		goto L3
	} else {
		goto L322
	}
L8:
	;
	v743 = int32(0)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v745 = F_errsave_start(m, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L40
	} else {
		goto L314
	}
L9:
	;
	if v83 != int32(34) {
		goto L7
	} else {
		goto L312
	}
L10:
	;
	if v83 == int32(124) {
		goto L8
	} else {
		goto L311
	}
L11:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v157 == int32(58) {
		goto L272
	} else {
		goto L273
	}
L12:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v606 = v30 - v605
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v604 <= v606+v607 {
		goto L244
	} else {
		goto L245
	}
L13:
	;
	if v157 != int32(34) {
		goto L11
	} else {
		goto L242
	}
L14:
	;
	if v157 == int32(124) {
		goto L12
	} else {
		goto L241
	}
L15:
	;
	v568 = m.G0
	v570 = v568 - int32(16)
	m.G0 = v570
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v573 = F_errsave_start(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L40
	} else {
		goto L231
	}
L16:
	;
	if int32(1)<<(uint(v165)%32)&int32(134218145) == int32(0) {
		goto L14
	} else {
		goto L230
	}
L17:
	;
	if int32(1)<<(uint(v104)%32)&int32(134218145) == int32(0) {
		goto L10
	} else {
		goto L229
	}
L18:
	;
	v802 = int32(4)
	goto L4
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L40
	} else {
		goto L226
	}
L20:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	switch v447 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L193
	default:
		goto L192
	case 42, 65, 97:
		goto L197
	case 44:
		v804 = int32(6)
		v805 = v30
		v807 = v32
		v808 = v33
		v810 = v36
		v812 = v38
		goto L3
	case 66, 98:
		goto L196
	case 67, 99:
		goto L195
	case 68, 100:
		goto L194
	}
L21:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if base.Ui32((v305-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L145
	} else {
		goto L146
	}
L22:
	;
	if l3 != 0 {
		goto L133
	} else {
		goto L134
	}
L23:
	;
	if l3 != 0 {
		goto L120
	} else {
		goto L121
	}
L24:
	;
	v248 = int32(0)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v250 = F_errsave_start(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L40
	} else {
		goto L111
	}
L25:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v227 = v30 - v45
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v226 <= v227+v228 {
		goto L102
	} else {
		goto L103
	}
L26:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v170 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L27:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if base.B2i32(v153 == int32(0))&base.B2i32(v157 == int32(92)) != 0 {
		goto L5
	} else {
		goto L71
	}
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v83 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v78 != int32(58) {
		goto L22
	} else {
		goto L44
	}
L31:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v41 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = v30 - v50
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v52 <= v53+v54 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = v44
	goto L32
L34:
	;
	goto L35
L35:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v47 == int32(39) {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v50 = v45
	goto L32
L37:
	;
	v58 = v52 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v58
	v60 = F_repalloc(m, v50, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v67 = v30
	goto L39
L39:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 == v70 {
		goto L24
	} else {
		goto L42
	}
L40:
	;
	return int32(0)
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v60
	v67 = v60 + v53
	goto L39
L42:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v72 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v74 = v67
	goto L30
L44:
	;
	v804 = int32(6)
	v805 = v74
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L45:
	;
	v783 = int32(0)
	goto L6
L46:
	;
	goto L47
L47:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v83 != int32(39) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.B2i32(v87&int32(1) == int32(0))&base.B2i32(v83 == int32(92)) != 0 {
		goto L5
	} else {
		goto L51
	}
L49:
	;
	if v87&int32(1) != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v804 = int32(4)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L51:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v100 != int32(1) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v104 = v83 - int32(33)
	if base.Ui32(v104) <= base.Ui32(int32(27)) {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	goto L10
L54:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v113 = F_errsave_start(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L40
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = v30 - v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v132 <= v134+v135 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	if v113 == int32(0) {
		v783 = v111
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L40
	} else {
		goto L59
	}
L59:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v120
	F_errmsg(m, int32(683775), v19+int32(32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L40
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, v112, int32(472131), int32(221), int32(197833))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L40
	} else {
		goto L61
	}
L61:
	;
	v783 = v111
	goto L6
L62:
	;
	v139 = v132 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v139
	v141 = F_repalloc(m, v133, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L40
	} else {
		goto L65
	}
L63:
	;
	v146 = v107
	v147 = v30
	goto L64
L64:
	;
	v148 = F_pg_mblen_cstr(m, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L40
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v141
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v146 = v145
	v147 = v141 + v134
	goto L64
L66:
	;
	if v148 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v804 = v38
	v805 = v151 + v148
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L68:
	;
	v150 = F__emscripten_memcpy_bulkmem(m, v147, v146, v148)
	mBase = m.M
	v151 = v150
	goto L70
L69:
	;
	v151 = v147
	goto L70
L70:
	;
	goto L67
L71:
	;
	switch v157 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L12
	default:
		goto L72
	}
L72:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v161 != int32(1) {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v165 = v157 - int32(33)
	if base.Ui32(v165) <= base.Ui32(int32(27)) {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	goto L14
L75:
	;
	if v169 == int32(39) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v169 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v804 = int32(8)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L79:
	;
	goto L80
L80:
	;
	if v169 == int32(92) {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v180 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v182 = F_errsave_start(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L40
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v206 = v30 - v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v204 <= v206+v207 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	if v182 == int32(0) {
		v783 = v180
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L40
	} else {
		goto L87
	}
L87:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v190
	if v189 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v194 = int32(682753)
	goto L90
L89:
	;
	v194 = int32(683295)
	goto L90
L90:
	;
	F_errmsg(m, v194, v19+int32(80))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L40
	} else {
		goto L91
	}
L91:
	;
	F_errsave_finish(m, v181, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L40
	} else {
		goto L92
	}
L92:
	;
	v783 = v180
	goto L6
L93:
	;
	v211 = v204 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v211
	v213 = F_repalloc(m, v205, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L40
	} else {
		goto L96
	}
L94:
	;
	v218 = v168
	v219 = v30
	goto L95
L95:
	;
	v220 = F_pg_mblen_cstr(m, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L40
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v213
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = v217
	v219 = v213 + v206
	goto L95
L97:
	;
	if v220 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v804 = int32(4)
	v805 = v223 + v220
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L99:
	;
	v222 = F__emscripten_memcpy_bulkmem(m, v219, v218, v220)
	mBase = m.M
	v223 = v222
	goto L101
L100:
	;
	v223 = v219
	goto L101
L101:
	;
	goto L98
L102:
	;
	v232 = v226 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v232
	v234 = F_repalloc(m, v45, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L40
	} else {
		goto L105
	}
L103:
	;
	v240 = v30
	v241 = v46
	goto L104
L104:
	;
	v242 = F_pg_mblen_cstr(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L40
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = v234 + v227
	v241 = v238
	goto L104
L106:
	;
	if v242 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v804 = int32(4)
	v805 = v245 + v242
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L108:
	;
	v244 = F__emscripten_memcpy_bulkmem(m, v240, v241, v242)
	mBase = m.M
	v245 = v244
	goto L110
L109:
	;
	v245 = v240
	goto L110
L110:
	;
	goto L107
L111:
	;
	if v250 == int32(0) {
		v783 = v248
		goto L6
	} else {
		goto L112
	}
L112:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L40
	} else {
		goto L113
	}
L113:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v258
	if v257 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v262 = int32(682753)
	goto L116
L115:
	;
	v262 = int32(683295)
	goto L116
L116:
	;
	F_errmsg(m, v262, v19+int32(96))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L40
	} else {
		goto L117
	}
L117:
	;
	F_errsave_finish(m, v249, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L40
	} else {
		goto L118
	}
L118:
	;
	v783 = v248
	goto L6
L119:
	;
	if l1 != 0 {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32
	goto L119
L121:
	;
	goto L122
L122:
	;
	if v33 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	F_pfree(m, v33)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L40
	} else {
		goto L124
	}
L124:
	;
	goto L119
L125:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v278
	goto L127
L126:
	;
	goto L127
L127:
	;
	if l2 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v67 - v280
	goto L130
L129:
	;
	goto L130
L130:
	;
	v283 = int32(1)
	if l5 == int32(0) {
		v783 = v283
		goto L6
	} else {
		goto L131
	}
L131:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v286
	v783 = v283
	goto L6
L132:
	;
	if l1 != 0 {
		goto L138
	} else {
		goto L139
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v33 == int32(0) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	F_pfree(m, v33)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L40
	} else {
		goto L137
	}
L137:
	;
	goto L132
L138:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v294
	goto L140
L139:
	;
	goto L140
L140:
	;
	if l2 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v74 - v296
	goto L143
L142:
	;
	goto L143
L143:
	;
	v299 = int32(1)
	if l5 == int32(0) {
		v783 = v299
		goto L6
	} else {
		goto L144
	}
L144:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v302
	v783 = v299
	goto L6
L145:
	;
	if v36 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	v421 = int32(0)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v423 = F_errsave_start(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L40
	} else {
		goto L184
	}
L148:
	;
	v333 = v330 + v329<<(uint(int32(1))%32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v339 = v335
	goto L158
L149:
	;
	v317 = F_palloc(m, int32(8))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L40
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v32+int32(1) < v36 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v328 = int32(4)
	v329 = int32(0)
	v330 = v317
	goto L148
L153:
	;
	v328 = v36
	v329 = v32
	v330 = v33
	goto L148
L154:
	;
	goto L155
L155:
	;
	v326 = F_repalloc(m, v33, v36<<(uint(int32(2))%32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L40
	} else {
		goto L156
	}
L156:
	;
	v328 = v36 << (uint(int32(1)) % 32)
	v329 = v32
	v330 = v326
	goto L148
L157:
	;
	v384 = int32(16383)
	if v384 < v383 {
		goto L173
	} else {
		goto L174
	}
L158:
	;
	v344 = v339 + int32(1)
	v345 = int32(*(*int8)(unsafe.Add(mBase, uint32(v339))))
	v346 = F___isspace(m, v345)
	mBase = m.M
	if v346 != 0 {
		v339 = v344
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v347 = int32(1)
	switch v345&int32(255) - int32(43) {
	case 0:
		v353 = v347
		goto L162
	default:
		v355 = v345
		v356 = v339
		v357 = v347
		goto L161
	case 2:
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v358 = int32(0)
	v360 = v355 - int32(48)
	if base.Ui32(v360) <= base.Ui32(int32(9)) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v354 = int32(*(*int8)(unsafe.Add(mBase, uint32(v344))))
	v355 = v354
	v356 = v344
	v357 = v353
	goto L161
L163:
	;
	v353 = int32(0)
	goto L162
L164:
	;
	v363 = v358
	v364 = v360
	v365 = v356
	goto L167
L165:
	;
	v377 = v358
	goto L166
L166:
	;
	if v357 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v367 = int32(10)
	v369 = v363*v367 - v364
	v370 = int32(*(*int8)(unsafe.Add(mBase, uint32(v365)+1)))
	v374 = v370 - int32(48)
	if base.Ui32(v374) < base.Ui32(v367) {
		v363 = v369
		v364 = v374
		v365 = v365 + int32(1)
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v377 = v369
	goto L166
L169:
	;
	goto L168
L170:
	;
	v383 = int32(0) - v377
	goto L172
L171:
	;
	v383 = v377
	goto L172
L172:
	;
	goto L157
L173:
	;
	v388 = int32(16383)
	goto L175
L174:
	;
	v388 = v383 & v384
	goto L175
L175:
	;
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333))))
	v392 = v388 | v389&int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v333))) = uint16(v392)
	if v388 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v398 = F_errsave_start(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L40
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v333))) = uint16(v388)
	v804 = int32(7)
	v805 = v30
	v807 = v329 + int32(1)
	v808 = v330
	v810 = v328
	v812 = v38
	goto L3
L179:
	;
	if v398 == int32(0) {
		v783 = v396
		goto L6
	} else {
		goto L180
	}
L180:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L40
	} else {
		goto L181
	}
L181:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v405
	F_errmsg(m, int32(683326), v19+int32(112))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L40
	} else {
		goto L182
	}
L182:
	;
	F_errsave_finish(m, v397, int32(472131), int32(335), int32(197833))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L40
	} else {
		goto L183
	}
L183:
	;
	v783 = v396
	goto L6
L184:
	;
	if v423 == int32(0) {
		v783 = v421
		goto L6
	} else {
		goto L185
	}
L185:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L40
	} else {
		goto L186
	}
L186:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v431
	if v430 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v435 = int32(682753)
	goto L189
L188:
	;
	v435 = int32(683295)
	goto L189
L189:
	;
	F_errmsg(m, v435, v19+int32(128))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L40
	} else {
		goto L190
	}
L190:
	;
	F_errsave_finish(m, v422, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L40
	} else {
		goto L191
	}
L191:
	;
	v783 = v421
	goto L6
L192:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v447-int32(48))&int32(255)) {
		goto L15
	} else {
		goto L225
	}
L193:
	;
	if l3 != 0 {
		goto L213
	} else {
		goto L214
	}
L194:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(1))%32)-int32(2)))))
	if base.Ui32(int32(16384)) <= base.Ui32(v513) {
		goto L15
	} else {
		goto L211
	}
L195:
	;
	v500 = v33 + v32<<(uint(int32(1))%32) - int32(2)
	v501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v500))))
	if base.Ui32(int32(16384)) <= base.Ui32(v501) {
		goto L15
	} else {
		goto L210
	}
L196:
	;
	v488 = v33 + v32<<(uint(int32(1))%32) - int32(2)
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488))))
	if base.Ui32(int32(16384)) <= base.Ui32(v489) {
		goto L15
	} else {
		goto L209
	}
L197:
	;
	v452 = v33 + v32<<(uint(int32(1))%32) - int32(2)
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	if base.Ui32(int32(16384)) <= base.Ui32(v453) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v456 = int32(0)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v458 = F_errsave_start(m, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L40
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v481 = v453 | int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v481)
	v804 = int32(7)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L201:
	;
	if v458 == int32(0) {
		v783 = v456
		goto L6
	} else {
		goto L202
	}
L202:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L40
	} else {
		goto L203
	}
L203:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v466
	if v465 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v470 = int32(682753)
	goto L206
L205:
	;
	v470 = int32(683295)
	goto L206
L206:
	;
	F_errmsg(m, v470, v19+int32(144))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L40
	} else {
		goto L207
	}
L207:
	;
	F_errsave_finish(m, v457, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L40
	} else {
		goto L208
	}
L208:
	;
	v783 = v456
	goto L6
L209:
	;
	v493 = v489 | int32(32768)
	*(*uint16)(unsafe.Add(mBase, uint32(v488))) = uint16(v493)
	v804 = int32(7)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L210:
	;
	v505 = v501 | int32(16384)
	*(*uint16)(unsafe.Add(mBase, uint32(v500))) = uint16(v505)
	v804 = int32(7)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L211:
	;
	v804 = int32(7)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L212:
	;
	if l1 != 0 {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32
	goto L212
L214:
	;
	goto L215
L215:
	;
	if v33 == int32(0) {
		goto L212
	} else {
		goto L216
	}
L216:
	;
	F_pfree(m, v33)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L40
	} else {
		goto L217
	}
L217:
	;
	goto L212
L218:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v523
	goto L220
L219:
	;
	goto L220
L220:
	;
	if l2 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30 - v525
	goto L223
L222:
	;
	goto L223
L223:
	;
	v528 = int32(1)
	if l5 == int32(0) {
		v783 = v528
		goto L6
	} else {
		goto L224
	}
L224:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v531
	v783 = v528
	goto L6
L225:
	;
	v804 = int32(7)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v29
	F_errmsg_internal(m, int32(459567), v19)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L40
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(472131), int32(378), int32(197833))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L40
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	goto L8
L230:
	;
	goto L12
L231:
	;
	if v573 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L40
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	m.G0 = v570 + int32(16)
	v783 = int32(0)
	goto L6
L235:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v579
	if v578 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v583 = int32(682753)
	goto L238
L237:
	;
	v583 = int32(683295)
	goto L238
L238:
	;
	F_errmsg(m, v583, v570)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L40
	} else {
		goto L239
	}
L239:
	;
	F_errsave_finish(m, v572, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L40
	} else {
		goto L240
	}
L240:
	;
	goto L234
L241:
	;
	goto L13
L242:
	;
	if v153 == int32(0) {
		goto L11
	} else {
		goto L243
	}
L243:
	;
	goto L12
L244:
	;
	v611 = v604 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v611
	v613 = F_repalloc(m, v605, v611)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L40
	} else {
		goto L247
	}
L245:
	;
	v617 = v605
	v618 = v30
	goto L246
L246:
	;
	if v617 == v618 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v613
	v617 = v613
	v618 = v613 + v606
	goto L246
L248:
	;
	v620 = int32(0)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v622 = F_errsave_start(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L40
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v644 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v644)
	if l3 != 0 {
		goto L260
	} else {
		goto L261
	}
L251:
	;
	if v622 == int32(0) {
		v783 = v620
		goto L6
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L40
	} else {
		goto L253
	}
L253:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v630
	if v629 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v634 = int32(682753)
	goto L256
L255:
	;
	v634 = int32(683295)
	goto L256
L256:
	;
	F_errmsg(m, v634, v19+int32(48))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L40
	} else {
		goto L257
	}
L257:
	;
	F_errsave_finish(m, v621, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L40
	} else {
		goto L258
	}
L258:
	;
	v783 = v620
	goto L6
L259:
	;
	if l1 != 0 {
		goto L265
	} else {
		goto L266
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32
	goto L259
L261:
	;
	goto L262
L262:
	;
	if v33 == int32(0) {
		goto L259
	} else {
		goto L263
	}
L263:
	;
	F_pfree(m, v33)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L40
	} else {
		goto L264
	}
L264:
	;
	goto L259
L265:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v652
	goto L267
L266:
	;
	goto L267
L267:
	;
	if l2 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v618 - v654
	goto L270
L269:
	;
	goto L270
L270:
	;
	v657 = int32(1)
	if l5 == int32(0) {
		v783 = v657
		goto L6
	} else {
		goto L271
	}
L271:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v660
	v783 = v657
	goto L6
L272:
	;
	if v30 == v662 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v713 = v30 - v662
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v712 <= v713+v714 {
		goto L302
	} else {
		goto L303
	}
L275:
	;
	v666 = int32(0)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v668 = F_errsave_start(m, v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L40
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30))) = uint8(v690)
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v692 != int32(1) {
		goto L286
	} else {
		goto L287
	}
L278:
	;
	if v668 == int32(0) {
		v783 = v666
		goto L6
	} else {
		goto L279
	}
L279:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L40
	} else {
		goto L280
	}
L280:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v676
	if v675 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v680 = int32(682753)
	goto L283
L282:
	;
	v680 = int32(683295)
	goto L283
L283:
	;
	F_errmsg(m, v680, v19-int32(-64))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L40
	} else {
		goto L284
	}
L284:
	;
	F_errsave_finish(m, v667, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L40
	} else {
		goto L285
	}
L285:
	;
	v783 = v666
	goto L6
L286:
	;
	v804 = int32(6)
	v805 = v30
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L287:
	;
	goto L288
L288:
	;
	if l3 != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	if l1 != 0 {
		goto L295
	} else {
		goto L296
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v32
	goto L289
L291:
	;
	goto L292
L292:
	;
	if v33 == int32(0) {
		goto L289
	} else {
		goto L293
	}
L293:
	;
	F_pfree(m, v33)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L40
	} else {
		goto L294
	}
L294:
	;
	goto L289
L295:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v702
	goto L297
L296:
	;
	goto L297
L297:
	;
	if l2 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30 - v704
	goto L300
L299:
	;
	goto L300
L300:
	;
	v707 = int32(1)
	if l5 == int32(0) {
		v783 = v707
		goto L6
	} else {
		goto L301
	}
L301:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v710
	v783 = v707
	goto L6
L302:
	;
	v718 = v712 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v718
	v720 = F_repalloc(m, v662, v718)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L40
	} else {
		goto L305
	}
L303:
	;
	v725 = v30
	v726 = v156
	goto L304
L304:
	;
	v727 = F_pg_mblen_cstr(m, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L40
	} else {
		goto L306
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v720
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v725 = v713 + v720
	v726 = v724
	goto L304
L306:
	;
	if v727 != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v804 = int32(2)
	v805 = v730 + v727
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L308:
	;
	v729 = F__emscripten_memcpy_bulkmem(m, v725, v726, v727)
	mBase = m.M
	v730 = v729
	goto L310
L309:
	;
	v730 = v725
	goto L310
L310:
	;
	goto L307
L311:
	;
	goto L9
L312:
	;
	if v87&int32(1) == int32(0) {
		goto L7
	} else {
		goto L313
	}
L313:
	;
	goto L8
L314:
	;
	if v745 == int32(0) {
		v783 = v743
		goto L6
	} else {
		goto L315
	}
L315:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L40
	} else {
		goto L316
	}
L316:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v753
	if v752 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v757 = int32(682753)
	goto L319
L318:
	;
	v757 = int32(683295)
	goto L319
L319:
	;
	F_errmsg(m, v757, v19+int32(16))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L40
	} else {
		goto L320
	}
L320:
	;
	F_errsave_finish(m, v744, int32(472131), int32(148), int32(199762))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L40
	} else {
		goto L321
	}
L321:
	;
	v783 = v743
	goto L6
L322:
	;
	if v83 == int32(32) {
		v804 = v767
		v805 = v30
		v807 = v32
		v808 = v33
		v810 = v36
		v812 = v38
		goto L3
	} else {
		goto L323
	}
L323:
	;
	v774 = F_pg_mblen_cstr(m, v82)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L40
	} else {
		goto L324
	}
L324:
	;
	if v774 != 0 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v804 = int32(2)
	v805 = v777 + v774
	v807 = v32
	v808 = v33
	v810 = v36
	v812 = v38
	goto L3
L326:
	;
	v776 = F__emscripten_memcpy_bulkmem(m, v30, v82, v774)
	mBase = m.M
	v777 = v776
	goto L328
L327:
	;
	v777 = v30
	goto L328
L328:
	;
	goto L325
L329:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v814 + v816
	v29 = v804
	v30 = v805
	v32 = v807
	v33 = v808
	v36 = v810
	v38 = v812
	goto L1
}
func F_ginadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v13 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v14 <= v13 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v13
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	v29 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+24)) = uint16(v29)
	v32 = v18 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 < v33 {
		v18 = v32
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(16)
	return
L8:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v44 <= v43 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = v43
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v48<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui32(int32(7)) < base.Ui32(v58) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L7
L12:
	;
	v95 = v48 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v95 < v96 {
		v48 = v95
		goto L10
	} else {
		goto L23
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+28)) = l0
	v92 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+24)) = uint16(v92)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v62 = int32(1) << (uint(v58) % 32)
	if v62&int32(242) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v62&int32(12) == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+24)) = uint8(v69)
	goto L12
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v79
	F_errmsg(m, int32(186020), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(475208), int32(324), int32(126378))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	goto L11
}
func F_ginbuildphasename(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = l0 - int64(1)
	if base.Ui64(v4) <= base.Ui64(int64(5)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)<<(uint(int32(2))%32))+uint32(_consts[48])))
		v13 = v12
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_ginendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+uint32(_consts[47])))
			F_MemoryContextDelete(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, v2)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ginrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ginFreeScanKeys(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if l1 == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v11 <= int32(0) {
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v16 = v11 * int32(48)
				if v16 != 0 {
					v17 = F__emscripten_memcpy_bulkmem(m, v14, l1, v16)
					mBase = m.M
				} else {
				}
			}
		}
		return
	}
}
func F_ginvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v586 int64
	_ = v586
	var v587 int64
	_ = v587
	var v592 int32
	_ = v592
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int64
	_ = v629
	var v633 int64
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	v23 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if int32(0) < v300 {
		goto L62
	} else {
		goto L63
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v33 = F_get_opfamily_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L2
	} else {
		goto L59
	}
L7:
	;
	v37 = int32(0)
	v39 = F_SearchSysCacheList(m, int32(4), int32(1), v32, v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(1)
	v44 = int32(0)
	v46 = F_SearchSysCacheList(m, int32(5), v41, v32, v44, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v48 <= int32(0) {
		v288 = v41
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v31
	goto L13
L12:
	;
	v51 = v30
	goto L13
L13:
	;
	v61 = int32(0)
	v64 = v41
	goto L14
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v61<<(uint(int32(2))%32))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+56))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
	v82 = v80 + v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	if v83 == v84 {
		v113 = v64
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v288 = v264
	goto L1
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if v114 != v30 {
		v264 = v113
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v86 = int32(0)
	v89 = F_errstart(m, int32(17), v86)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v89 == int32(0) {
		v113 = v86
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v97 = F_format_procedure(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v33
	F_errmsg(m, int32(152080), v20+int32(256))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(475208), int32(85), int32(339378))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v113 = v86
	goto L16
L24:
	;
	v267 = v61 + int32(1)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v267 < v268 {
		v61 = v267
		v64 = v264
		goto L14
	} else {
		goto L58
	}
L25:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+16)))
	switch v116 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L35
	case 2:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L32
	case 5:
		goto L31
	case 6:
		goto L30
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L54
	}
L27:
	;
	v230 = int32(0)
	v233 = F_errstart(m, int32(17), v230)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L2
	} else {
		goto L52
	}
L28:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v51
	v223 = int32(2)
	v227 = F_check_amproc_signature(m, v218, int32(23), int32(0), v223, v223, v20+int32(112))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L50
	}
L29:
	;
	v209 = int32(0)
	v212 = F_errstart(m, int32(17), v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L48
	}
L30:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v205 = F_check_amoptsproc_signature(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L46
	}
L31:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+244)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+236)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v20)+224)) = int64(90194315497)
	v196 = int32(7)
	v200 = F_check_amproc_signature(m, v186, int32(18), int32(0), v196, v196, v20+int32(224))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = int64(9796820402197)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v51
	v178 = int32(4)
	v182 = F_check_amproc_signature(m, v171, int32(23), int32(0), v178, v178, v20+int32(208))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L42
	}
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(204)))) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+196)) = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+188)) = int64(9796820402199)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v20)+176)) = int64(90194315497)
	v167 = F_check_amproc_signature(m, v151, int32(16), int32(0), int32(6), int32(8), v20+int32(176))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v134 = int64(9796820404457)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(164)))) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v20)+156)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v20)+148)) = int64(90194315497)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v30
	v147 = F_check_amproc_signature(m, v133, int32(2281), int32(0), int32(5), int32(7), v20+int32(144))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+132)) = int64(9796820404457)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v30
	v129 = F_check_amproc_signature(m, v119, int32(2281), int32(0), int32(2), int32(3), v20+int32(128))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v129 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v264 = v113
	goto L24
L38:
	;
	if v147 == int32(0) {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v264 = v113
	goto L24
L40:
	;
	if v167 == int32(0) {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v264 = v113
	goto L24
L42:
	;
	if v182 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v264 = v113
	goto L24
L44:
	;
	if v200 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v264 = v113
	goto L24
L46:
	;
	if v205 == int32(0) {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v264 = v113
	goto L24
L48:
	;
	if v212 == int32(0) {
		v264 = v209
		goto L24
	} else {
		goto L49
	}
L49:
	;
	v240 = int32(145)
	v241 = int32(449639)
	goto L26
L50:
	;
	if v227 != 0 {
		v264 = v113
		goto L24
	} else {
		goto L51
	}
L51:
	;
	goto L27
L52:
	;
	if v233 == int32(0) {
		v264 = v230
		goto L24
	} else {
		goto L53
	}
L53:
	;
	v240 = int32(157)
	v241 = int32(449534)
	goto L26
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v246 = F_format_procedure(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v20)+104)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v33
	F_errmsg(m, v241, v20+int32(96))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(475208), v240, int32(339378))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v264 = int32(0)
	goto L24
L58:
	;
	goto L15
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l0
	F_errmsg_internal(m, int32(39938), v20)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(475208), int32(51), int32(339378))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v308 = int32(0)
	v311 = v288
	goto L65
L63:
	;
	v445 = v288
	goto L64
L64:
	;
	v457 = F_identify_opfamily_groups(m, v39, v46)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L2
	} else {
		goto L97
	}
L65:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(48)+v308<<(uint(int32(2))%32))))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+56))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+22)))
	v329 = v327 + v328
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v329)+16)))
	if base.Ui32(int32(65472)) < base.Ui32((v330+int32(-64))&int32(65535)) {
		v366 = v311
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v445 = v435
	goto L64
L67:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+18)))
	if v368 == int32(115) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v337 = int32(0)
	v340 = F_errstart(m, int32(17), v337)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	if v340 == int32(0) {
		v366 = v337
		goto L67
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	v348 = F_format_operator(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v329)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v33
	F_errmsg(m, int32(449410), v20+int32(80))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(475208), int32(176), int32(339378))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v366 = v337
	goto L67
L75:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v406 = F_check_amop_signature(m, v402, int32(16), v404, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L87
	}
L76:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v329)+28))
	if v371 == int32(0) {
		v401 = v366
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v374 = int32(0)
	v377 = F_errstart(m, int32(17), v374)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L2
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	if v377 == int32(0) {
		v401 = v374
		goto L75
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	v385 = F_format_operator(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v33
	F_errmsg(m, int32(170703), v20-int32(-64))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(475208), int32(188), int32(339378))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v401 = v374
	goto L75
L86:
	;
	v437 = v308 + int32(1)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	if v437 < v438 {
		v308 = v437
		v311 = v435
		goto L65
	} else {
		goto L95
	}
L87:
	;
	if v406 != 0 {
		v435 = v401
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v408 = int32(0)
	v411 = F_errstart(m, int32(17), v408)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v411 == int32(0) {
		v435 = v408
		goto L86
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v329)+20))
	v419 = F_format_operator(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v33
	F_errmsg(m, int32(345520), v20+int32(48))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(475208), int32(201), int32(339378))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v435 = v408
	goto L86
L95:
	;
	goto L66
L96:
	;
	v568 = v29 + int32(8)
	v575 = v445
	v586 = int64(1)
	goto L131
L97:
	;
	if v457 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v550 = int32(0)
	goto L96
L99:
	;
	goto L100
L100:
	;
	v462 = int32(0)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	if v463 <= v462 {
		v550 = v462
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v466 = int32(0)
	if v466 < v463 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v469 = v463
	goto L104
L103:
	;
	v469 = v466
	goto L104
L104:
	;
	v470 = int32(1)
	if v463 == v470 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v469&v470 == int32(0) {
		v550 = v521
		goto L96
	} else {
		goto L124
	}
L106:
	;
	v474 = int32(0)
	v521 = v474
	v523 = v474
	goto L105
L107:
	;
	goto L108
L108:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	v479 = int32(0)
	v482 = v479
	v484 = v479
	v489 = v479
	goto L109
L109:
	;
	v501 = v478 + v484<<(uint(int32(2))%32)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v30 == v503 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v521 = v515
	v523 = v517
	goto L105
L111:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v502)+4))
	if v505 == v30 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v508 = v482
	goto L113
L113:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	if v30 == v510 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v507 = v502
	goto L116
L115:
	;
	v507 = v482
	goto L116
L116:
	;
	v508 = v507
	goto L113
L117:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v512 == v30 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v515 = v508
	goto L119
L119:
	;
	v516 = int32(2)
	v517 = v484 + v516
	v519 = v489 + v516
	if v519 != v469&int32(2147483646) {
		v482 = v515
		v484 = v517
		v489 = v519
		goto L109
	} else {
		goto L123
	}
L120:
	;
	v514 = v509
	goto L122
L121:
	;
	v514 = v508
	goto L122
L122:
	;
	v515 = v514
	goto L119
L123:
	;
	goto L110
L124:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v457)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v523<<(uint(int32(2))%32))))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)))
	if v545 != v30 {
		v550 = v521
		goto L96
	} else {
		goto L125
	}
L125:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v547 == v30 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v549 = v544
	goto L128
L127:
	;
	v549 = v521
	goto L128
L128:
	;
	v550 = v549
	goto L96
L129:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L2
	} else {
		goto L157
	}
L130:
	;
	v645 = int32(0)
	v648 = F_errstart(m, int32(17), v645)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L2
	} else {
		goto L152
	}
L131:
	;
	if v550 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+16)))
	if v639&int32(80) != 0 {
		v671 = v637
		goto L129
	} else {
		goto L151
	}
L133:
	;
	goto L132
L134:
	;
	v633 = v586 + int64(1)
	if v633 != int64(8) {
		v586 = v633
		goto L131
	} else {
		goto L150
	}
L135:
	;
	v587 = *(*int64)(unsafe.Add(mBase, uint32(v550)+16))
	if base.I32_wrap_i64(int64(base.Ui64(v587)>>(uint(v586)%64)))&int32(1) != 0 {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v592 = base.I32_wrap_i64(v586)
	if v592&int32(5) == int32(4) {
		v627 = v575
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	v629 = v586 + int64(1)
	if v629 != int64(8) {
		v575 = v627
		v586 = v629
		goto L131
	} else {
		goto L148
	}
L140:
	;
	if v586 == int64(7) {
		v627 = v575
		goto L139
	} else {
		goto L141
	}
L141:
	;
	if v592&int32(3) == int32(1) {
		v627 = v575
		goto L139
	} else {
		goto L142
	}
L142:
	;
	v603 = int32(0)
	v606 = F_errstart(m, int32(17), v603)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v606 == int32(0) {
		v627 = v603
		goto L139
	} else {
		goto L144
	}
L144:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v568
	F_errmsg(m, int32(450554), v20+int32(32))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(475208), int32(242), int32(339378))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L2
	} else {
		goto L147
	}
L147:
	;
	v627 = v603
	goto L139
L148:
	;
	if v550 != 0 {
		v637 = v627
		goto L133
	} else {
		goto L149
	}
L149:
	;
	goto L130
L150:
	;
	v637 = v575
	goto L133
L151:
	;
	goto L130
L152:
	;
	if v648 == int32(0) {
		v671 = v645
		goto L129
	} else {
		goto L153
	}
L153:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(25769803780)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(263377)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v568
	F_errmsg(m, int32(448999), v20+int32(16))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(475208), int32(253), int32(339378))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	v671 = v645
	goto L129
L157:
	;
	F_ReleaseCatCacheList(m, v39)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L2
	} else {
		goto L158
	}
L158:
	;
	F_ReleaseCatCache(m, v23)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L2
	} else {
		goto L159
	}
L159:
	;
	m.G0 = v20 + int32(272)
	return v671 & int32(1)
}
func F_gistbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if l1 == int32(0) {
		v8 = F_palloc0(m, int32(40))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = v8
			F_gistvacuumscan(m, l0, v12, l2, l3)
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	} else {
		v12 = l1
		F_gistvacuumscan(m, l0, v12, l2, l3)
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_gistcheckpage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l1^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v26 = v20 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
	if v27 != 0 {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
		if (v28<<(uint(int32(8))%32)-v31)&int32(65535) != int32(16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l1 < int32(0) {
						v94 = *(*int32)(unsafe.Add(mBase, _consts[8]))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(l1^int32(-1))<<(uint(int32(6))%32))+16))
						v109 = v100
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+l1<<(uint(int32(6))%32)+int32(-64))+16))
						v109 = v108
					}
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v109
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v90 + int32(4)
					F_errmsg(m, int32(45277), v7+int32(16))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						F_errhint(m, int32(541506), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							F_errfinish(m, int32(474055), int32(812), int32(387683))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
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
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if l1 < int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l1^int32(-1))<<(uint(int32(6))%32))+16))
					v66 = v57
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l1<<(uint(int32(6))%32)+int32(-64))+16))
					v66 = v65
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v66
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v47 + int32(4)
				F_errmsg(m, int32(45224), v7)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errhint(m, int32(541506), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errfinish(m, int32(474055), int32(801), int32(387683))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
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
func F_gistfillitupvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if v4 < l1 {
		if l1 != int32(1) {
			v20 = v4
			v21 = v4
			v22 = v4
			for {
				v25 = int32(2)
				v27 = l0 + v21<<(uint(v25)%32)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
				v30 = int32(8191)
				v32 = v20 + v29&v30
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
				v38 = v32 + v35&v30
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38
				v41 = v21 + v25
				v43 = v22 + v25
				if v43 != l1&int32(2147483646) {
					v20 = v38
					v21 = v41
					v22 = v43
					continue
				} else {
					break
				}
				break
			}
			v48 = v38
			v49 = v41
		} else {
			v48 = v4
			v49 = v4
		}
		if l1&int32(1) != 0 {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0+v49<<(uint(int32(2))%32))))
			v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+6)))
			v62 = v48 + v59&int32(8191)
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v62
			v64 = v62
		} else {
			v64 = v48
		}
		v66 = F_palloc(m, v64)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			v73 = v66
			v74 = int32(0)
			for {
				v80 = l0 + v74<<(uint(int32(2))%32)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+6)))
				v84 = v82 & int32(8191)
				if v84 != 0 {
					v85 = F__emscripten_memcpy_bulkmem(m, v73, v81, v84)
					mBase = m.M
					v86 = v85
				} else {
					v86 = v73
				}
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
				v93 = v74 + int32(1)
				if v93 != l1 {
					v73 = v86 + v88&int32(8191)
					v74 = v93
					continue
				} else {
					break
				}
				break
			}
			return v66
		}
	} else {
		v97 = F_palloc(m, int32(0))
		mBase = m.M
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			return v97
		}
	}
}
func F_gistfinishsplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v36 int32
	_ = v36
	var __phi36 int32
	_ = __phi36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_LockBuffer(m, v17, int32(2))
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
	if l3 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistFindCorrectParent(m, v88, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v25 = v23 - int32(1)
	if v25 < int32(2) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	__phi33 = v23
	__phi36 = v25
	v33 = __phi33
	v36 = __phi36
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v40 = int32(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+v33<<(uint(v40)%32)-int32(8))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39+v36<<(uint(v40)%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_gistFindCorrectParent(m, v50, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v62 = F_gistinserttuples(m, l0, v53, l2, v49+int32(4), int32(1), v57, v58, v59, v57, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v62 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v64 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v64)
	goto L12
L11:
	;
	goto L12
L12:
	;
	if int32(2) < v36 {
		__phi33 = v36
		__phi36 = v36 - int32(1)
		v33 = __phi33
		v36 = __phi36
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v99 = F_gistinserttuples(m, l0, v91, l2, v14+int32(8), int32(2), v95, v96, v97, int32(1), l4)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v101)
	v103 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v103)
	m.G0 = v14 + int32(16)
	return
}
func F_gistnospace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v6 = int32(0)
	if l2 <= v6 {
		v71 = l4
	} else {
		v11 = int32(1)
		if l2 == v11 {
			v49 = int32(0)
			v51 = l4
		} else {
			v21 = int32(0)
			v23 = l4
			v24 = v6
			for {
				v27 = int32(2)
				v29 = l1 + v21<<(uint(v27)%32)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+6)))
				v32 = int32(8191)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)))
				v41 = v23 + v31&v32 + v36&v32 + int32(8)
				v43 = v21 + v27
				v45 = v24 + v27
				if v45 != l2&int32(2147483646) {
					v21 = v43
					v23 = v41
					v24 = v45
					continue
				} else {
					break
				}
				break
			}
			v49 = v43
			v51 = v41
		}
		if l2&v11 == int32(0) {
			v71 = v51
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l1+v49<<(uint(int32(2))%32))))
			v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)))
			v71 = v61&int32(8191) + v51 + int32(4)
		}
	}
	if l3 != 0 {
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32)+l0)+20))
		v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v78&int32(32767))+6)))
		v88 = v82&int32(8191) + int32(4)
	} else {
		v88 = int32(0)
	}
	v89 = int32(4)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v92 = v90 - v91
	if v92 <= v89 {
		v95 = v89
	} else {
		v95 = v92
	}
	return base.B2i32(base.Ui32(v95-int32(4)+v88) < base.Ui32(v71))
}
func F_gistplacetopage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
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
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int64
	_ = v415
	var v416 int32
	_ = v416
	var v420 int64
	_ = v420
	var v421 int32
	_ = v421
	var v427 int64
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v796 int32
	_ = v796
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int64
	_ = v822
	var v823 int64
	_ = v823
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int64
	_ = v876
	var v877 int32
	_ = v877
	var v893 int32
	_ = v893
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1113 int32
	_ = v1113
	var v1114 int64
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int64
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int64
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1193 int32
	_ = v1193
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1230 int32
	_ = v1230
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1398 int32
	_ = v1398
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1428 int32
	_ = v1428
	var v1469 int32
	_ = v1469
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1681 int32
	_ = v1681
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1928 int64
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1933 int64
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1961 int64
	_ = v1961
	var v1981 int32
	_ = v1981
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2024 int32
	_ = v2024
	var v2040 int32
	_ = v2040
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2098 int32
	_ = v2098
	var v2108 int64
	_ = v2108
	var v2113 int32
	_ = v2113
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	v7 = l6
	v11 = l10
	v27 = m.G0
	v29 = v27 - int32(864)
	m.G0 = v29
	if l3 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v49 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v49 = v48
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L61
	} else {
		goto L387
	}
L6:
	;
	if l8 != 0 {
		goto L380
	} else {
		goto L381
	}
L7:
	;
	if v752 != 0 {
		goto L243
	} else {
		goto L244
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L61
	} else {
		goto L237
	}
L9:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v67)+12)))
	if v70&int32(8) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(l3^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v67 = v61 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v75
	if l5 <= v75 {
		v142 = l1
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L61
	} else {
		goto L234
	}
L16:
	;
	v1045 = int32(4437236)
	v1047 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1048 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1047 + v1048
	v1052 = v7 - v1048
	if base.Ui32(v1052&int32(65535)) <= base.Ui32(int32(2047)) {
		goto L201
	} else {
		goto L202
	}
L17:
	;
	if base.B2i32(base.Ui32(v160+v159) < base.Ui32(v142)) == int32(0) {
		goto L16
	} else {
		goto L31
	}
L18:
	;
	if v7 != 0 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v82 = int32(1)
	if l5 == v82 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if l5&v82 == int32(0) {
		v142 = v122
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v120 = int32(0)
	v122 = l1
	goto L20
L22:
	;
	goto L23
L23:
	;
	v92 = int32(0)
	v94 = l1
	v95 = v75
	goto L24
L24:
	;
	v98 = int32(2)
	v100 = l4 + v92<<(uint(v98)%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+6)))
	v103 = int32(8191)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+6)))
	v112 = v94 + v102&v103 + v107&v103 + int32(8)
	v114 = v92 + v98
	v116 = v95 + v98
	if v116 != l5&int32(2147483646) {
		v92 = v114
		v94 = v112
		v95 = v116
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v120 = v114
	v122 = v112
	goto L20
L26:
	;
	goto L25
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l4+v120<<(uint(int32(2))%32))))
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+6)))
	v142 = v132&int32(8191) + v122 + int32(4)
	goto L18
L28:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v7<<(uint(int32(2))%32)+v67)+20))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v149&int32(32767))+6)))
	v159 = v153&int32(8191) + int32(4)
	goto L30
L29:
	;
	v159 = int32(0)
	goto L30
L30:
	;
	v160 = F_PageGetFreeSpace(m, v67)
	mBase = m.M
	goto L17
L31:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v165)+12)))
	v168 = int32(17)
	if v167&v168 == v168 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if base.Ui32(v172) < base.Ui32(int32(25)) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v579 = F_gistextractpage(m, v67, v29+int32(44))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L61
	} else {
		goto L101
	}
L35:
	;
	v463 = int32(0)
	if l5 <= v463 {
		v528 = l1
		goto L87
	} else {
		goto L88
	}
L36:
	;
	v178 = int32(base.Ui32(v172+int32(262120)) >> (uint(int32(2)) % 32))
	if v178&int32(65535) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v183 = int32(1)
	v184 = int32(2)
	v188 = (v178 + v183) & int32(65535)
	if base.Ui32(v188) <= base.Ui32(v184) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v191 = v184
	goto L40
L39:
	;
	v191 = v188
	goto L40
L40:
	;
	v192 = int32(1)
	v193 = v191 - v192
	v197 = v67 + int32(24)
	if base.Ui32(v188) < base.Ui32(int32(3)) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v193&v192 == int32(0) {
		v317 = v285
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v284 = v183
	v285 = int32(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v203 = int32(0)
	v218 = v183
	v219 = v203
	v222 = v203
	goto L45
L45:
	;
	v233 = v218<<(uint(int32(2))%32) + v197
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233-int32(4))))
	v237 = int32(98304)
	if v236&v237 == v237 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v284 = v267
	v285 = v265
	goto L41
L47:
	;
	v243 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v219<<(uint(v243)%32)))) = uint16(v218)
	v249 = v219 + v243
	goto L49
L48:
	;
	v249 = v219
	goto L49
L49:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v251 = int32(98304)
	if v250&v251 == v251 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v257 = int32(1)
	v261 = v218 + v257
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v249<<(uint(v257)%32)))) = uint16(v261)
	v265 = v249 + v257
	goto L52
L51:
	;
	v265 = v249
	goto L52
L52:
	;
	v266 = int32(2)
	v267 = v218 + v266
	v269 = v222 + v266
	if v269 != v193&int32(-2) {
		v218 = v267
		v219 = v265
		v222 = v269
		goto L45
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	if v317 <= int32(0) {
		goto L35
	} else {
		goto L57
	}
L55:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v284<<(uint(int32(2))%32)+v197-int32(4))))
	v305 = int32(98304)
	if v304&v305 != v305 {
		v317 = v285
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v311 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+int32(48)+v285<<(uint(v311)%32)))) = uint16(v284)
	v317 = v285 + v311
	goto L54
L57:
	;
	v320 = int32(0)
	v322 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v322 <= v320 {
		v336 = v320
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v337 = int32(4437236)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v339 + int32(1)
	F_PageIndexMultiDelete(m, v67, v29+int32(48), v317)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L61
	} else {
		goto L63
	}
L59:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+118)))
	if v327 != int32(112) {
		v336 = int32(0)
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v332 = F_index_compute_xid_horizon_for_tuples(m, l0, l11, l3, v29+int32(48), v317)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	return int32(0)
L62:
	;
	v336 = v332
	goto L58
L63:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v348 = v67 + v347
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+12)))
	v351 = v349 & int32(65519)
	*(*uint16)(unsafe.Add(mBase, uint32(v348)+12)) = uint16(v351)
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+118)))
	if v356 != int32(112) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotr(v427, int64(32))
	v431 = int32(4437236)
	v433 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v433 - int32(1)
	goto L35
L66:
	;
	v420 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L61
	} else {
		goto L85
	}
L67:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v360 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v363 != 0 {
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v367 = int32(0)
	v368 = m.G0
	v370 = v368 - int32(16)
	m.G0 = v370
	v373 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v373 < int32(2) {
		v393 = v367
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v364 != 0 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+8)) = v336
	*(*uint8)(unsafe.Add(mBase, uint32(v370)+14)) = uint8(v393)
	*(*uint16)(unsafe.Add(mBase, uint32(v370)+12)) = uint16(v317)
	F_XLogBeginInsert(m)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L61
	} else {
		goto L80
	}
L74:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+118)))
	if v377 != int32(112) {
		v393 = v367
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l11)+56))
	goto L76
L76:
	;
	if base.Ui32(v381) < base.Ui32(int32(12000)) {
		v393 = int32(1)
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v384 = int32(0)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l11)+180))
	if v385 == v384 {
		v393 = v384
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l11)+48))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+119)))
	switch v389 - int32(109) {
	case 0, 5:
		goto L79
	default:
		v393 = v384
		goto L73
	}
L79:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+104)))
	v393 = v392
	goto L73
L80:
	;
	v400 = int32(8)
	F_XLogRegisterData(m, v370+v400, v400)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L61
	} else {
		goto L81
	}
L81:
	;
	F_XLogRegisterData(m, v29+int32(48), v317<<(uint(int32(1))%32))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L61
	} else {
		goto L82
	}
L82:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L61
	} else {
		goto L83
	}
L83:
	;
	v415 = F_XLogInsert(m, int32(14), int32(16))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L61
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v370 + int32(16)
	v427 = v415
	goto L65
L85:
	;
	v427 = v420
	goto L65
L86:
	;
	if base.B2i32(base.Ui32(v546+v545) < base.Ui32(v528)) == int32(0) {
		goto L16
	} else {
		goto L100
	}
L87:
	;
	if v7 != 0 {
		goto L97
	} else {
		goto L98
	}
L88:
	;
	v468 = int32(1)
	if l5 == v468 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if l5&v468 == int32(0) {
		v528 = v508
		goto L87
	} else {
		goto L96
	}
L90:
	;
	v506 = int32(0)
	v508 = l1
	goto L89
L91:
	;
	goto L92
L92:
	;
	v478 = int32(0)
	v480 = l1
	v481 = v463
	goto L93
L93:
	;
	v484 = int32(2)
	v486 = l4 + v478<<(uint(v484)%32)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487)+6)))
	v489 = int32(8191)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v492)+6)))
	v498 = v480 + v488&v489 + v493&v489 + int32(8)
	v500 = v478 + v484
	v502 = v481 + v484
	if v502 != l5&int32(2147483646) {
		v478 = v500
		v480 = v498
		v481 = v502
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v506 = v500
	v508 = v498
	goto L89
L95:
	;
	goto L94
L96:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l4+v506<<(uint(int32(2))%32))))
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517)+6)))
	v528 = v518&int32(8191) + v508 + int32(4)
	goto L87
L97:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v7<<(uint(int32(2))%32)+v67)+20))
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v535&int32(32767))+6)))
	v545 = v539&int32(8191) + int32(4)
	goto L99
L98:
	;
	v545 = int32(0)
	goto L99
L99:
	;
	v546 = F_PageGetFreeSpace(m, v67)
	mBase = m.M
	goto L86
L100:
	;
	goto L34
L101:
	;
	if base.Ui32(int32(2047)) < base.Ui32((v7-int32(1))&int32(65535)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v746 = int32(0)
	v749 = F_gistjoinvector(m, v579, v29+int32(44), l4, l5)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L61
	} else {
		goto L151
	}
L103:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v587 - int32(1)
	if v7 == v587 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v592 = int32(2)
	v594 = v579 + v7<<(uint(v592)%32)
	v596 = v594 - int32(4)
	v599 = (v587 - v7) << (uint(v592) % 32)
	if v596 == v594 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L102
L106:
	;
	goto L105
L107:
	;
	v603 = v596 + v599
	if base.Ui32(v594-v603) <= base.Ui32(int32(0)-v599<<(uint(int32(1))%32)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v610 = F___memcpy(m, v596, v594, v599)
	mBase = m.M
	goto L105
L109:
	;
	goto L110
L110:
	;
	v613 = (v596 ^ v594) & int32(3)
	if base.Ui32(v596) < base.Ui32(v594) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if v715 == int32(0) {
		goto L106
	} else {
		goto L147
	}
L112:
	;
	if base.Ui32(v693) <= base.Ui32(int32(3)) {
		v714 = v692
		v715 = v693
		v716 = v694
		goto L111
	} else {
		goto L143
	}
L113:
	;
	if v613 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if v613 != 0 {
		v675 = v599
		goto L126
	} else {
		goto L127
	}
L116:
	;
	v714 = v594
	v715 = v599
	v716 = v596
	goto L111
L117:
	;
	goto L118
L118:
	;
	if v596&int32(3) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v692 = v594
	v693 = v599
	v694 = v596
	goto L112
L120:
	;
	goto L121
L121:
	;
	v620 = v594
	v621 = v599
	v622 = v596
	goto L122
L122:
	;
	if v621 == int32(0) {
		goto L106
	} else {
		goto L124
	}
L123:
	;
	v692 = v629
	v693 = v631
	v694 = v633
	goto L112
L124:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	*(*uint8)(unsafe.Add(mBase, uint32(v622))) = uint8(v626)
	v628 = int32(1)
	v629 = v620 + v628
	v631 = v621 - v628
	v633 = v622 + v628
	if v633&int32(3) != 0 {
		v620 = v629
		v621 = v631
		v622 = v633
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	if v675 == int32(0) {
		goto L106
	} else {
		goto L139
	}
L127:
	;
	if v603&int32(3) != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v640 = v599
	goto L131
L129:
	;
	v655 = v599
	goto L130
L130:
	;
	if base.Ui32(v655) <= base.Ui32(int32(3)) {
		v675 = v655
		goto L126
	} else {
		goto L135
	}
L131:
	;
	if v640 == int32(0) {
		goto L106
	} else {
		goto L133
	}
L132:
	;
	v655 = v646
	goto L130
L133:
	;
	v646 = v640 - int32(1)
	v647 = v596 + v646
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v646))))
	*(*uint8)(unsafe.Add(mBase, uint32(v647))) = uint8(v649)
	if v647&int32(3) != 0 {
		v640 = v646
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v662 = v655
	goto L136
L136:
	;
	v666 = v662 - int32(4)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v594+v666)))
	*(*int32)(unsafe.Add(mBase, uint32(v596+v666))) = v669
	if base.Ui32(int32(3)) < base.Ui32(v666) {
		v662 = v666
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v675 = v666
	goto L126
L138:
	;
	goto L137
L139:
	;
	v682 = v675
	goto L140
L140:
	;
	v686 = v682 - int32(1)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v686))))
	*(*uint8)(unsafe.Add(mBase, uint32(v596+v686))) = uint8(v689)
	if v686 != 0 {
		v682 = v686
		goto L140
	} else {
		goto L142
	}
L141:
	;
	goto L106
L142:
	;
	goto L141
L143:
	;
	v699 = v692
	v700 = v693
	v701 = v694
	goto L144
L144:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v703
	v705 = int32(4)
	v706 = v699 + v705
	v708 = v701 + v705
	v710 = v700 - v705
	if base.Ui32(int32(3)) < base.Ui32(v710) {
		v699 = v706
		v700 = v710
		v701 = v708
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v714 = v706
	v715 = v710
	v716 = v708
	goto L111
L146:
	;
	goto L145
L147:
	;
	v721 = v714
	v722 = v715
	v723 = v716
	goto L148
L148:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721))))
	*(*uint8)(unsafe.Add(mBase, uint32(v723))) = uint8(v725)
	v727 = int32(1)
	v732 = v722 - v727
	if v732 != 0 {
		v721 = v721 + v727
		v722 = v732
		v723 = v723 + v727
		goto L148
	} else {
		goto L150
	}
L149:
	;
	goto L106
L150:
	;
	goto L149
L151:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v752 = F_gistSplit(m, l0, v67, v749, v751, l2)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L61
	} else {
		goto L152
	}
L152:
	;
	if v752 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v767 = v746
	v768 = v752
	goto L156
L154:
	;
	v796 = v746
	goto L155
L155:
	;
	v811 = v796 + base.B2i32(v49 == int32(0))
	if int32(76) <= v811 {
		goto L8
	} else {
		goto L159
	}
L156:
	;
	v781 = v767 + int32(1)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v768)+28))
	if v782 != 0 {
		v767 = v781
		v768 = v782
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v796 = v781
	goto L155
L158:
	;
	goto L157
L159:
	;
	v815 = v70 & int32(1)
	if v49 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v877 == int32(0) {
		goto L7
	} else {
		goto L173
	}
L161:
	;
	v874 = int32(-1)
	v876 = int64(0)
	v877 = v752
	goto L160
L162:
	;
	goto L163
L163:
	;
	v819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v820 = v67 + v819
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)+8))
	v822 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v820)+4)))
	v823 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v820))))
	*(*int32)(unsafe.Add(mBase, uint32(v752)+24)) = l3
	if l3 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v752))) = v843
	if l3 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v828+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v843 = v834
	goto L164
L166:
	;
	goto L167
L167:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v836+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v843 = v842
	goto L164
L168:
	;
	v866 = F_PageGetTempPageCopySpecial(m, v865)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L61
	} else {
		goto L172
	}
L169:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v851+(l3^int32(-1))<<(uint(int32(2))%32))))
	v865 = v857
	goto L168
L170:
	;
	goto L171
L171:
	;
	v859 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v865 = v859 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L168
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v752)+20)) = v866
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v866)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v866+v869)+12)) = uint16(v815)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v752)+28))
	v874 = v821
	v876 = v823<<(uint(int64(32))%64) | v822
	v877 = v872
	goto L160
L173:
	;
	v893 = v877
	goto L174
L174:
	;
	v906 = F_gistNewBuffer(m, l0, l11)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L61
	} else {
		goto L176
	}
L175:
	;
	goto L7
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v893)+24)) = v906
	if v906 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v893)+24))
	if v937 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L178:
	;
	F_PageInit(m, v926, int32(8192), int32(16))
	mBase = m.M
	v930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v926)+16)))
	v931 = v926 + v930
	v932 = int32(65409)
	*(*uint16)(unsafe.Add(mBase, uint32(v931)+14)) = uint16(v932)
	*(*uint16)(unsafe.Add(mBase, uint32(v931)+12)) = uint16(v815)
	*(*int32)(unsafe.Add(mBase, uint32(v931)+8)) = int32(-1)
	goto L177
L179:
	;
	v912 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v912+(v906^int32(-1))<<(uint(int32(2))%32))))
	v926 = v918
	goto L178
L180:
	;
	goto L181
L181:
	;
	v920 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v926 = v920 + v906<<(uint(int32(13))%32) + int32(-8192)
	goto L178
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v893)+20)) = v955
	if v937 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v941+(v937^int32(-1))<<(uint(int32(2))%32))))
	v955 = v947
	goto L182
L184:
	;
	goto L185
L185:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v955 = v949 + v937<<(uint(int32(13))%32) + int32(-8192)
	goto L182
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v893))) = v975
	if l3 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L187:
	;
	v960 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v960+(v937^int32(-1))<<(uint(int32(6))%32))+16))
	v975 = v966
	goto L186
L188:
	;
	goto L189
L189:
	;
	v968 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v968+v937<<(uint(int32(6))%32)+int32(-64))+16))
	v975 = v974
	goto L186
L190:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v893)+24))
	if v996 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v980 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v980+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v995 = v986
	goto L190
L192:
	;
	goto L193
L193:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v988+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v995 = v994
	goto L190
L194:
	;
	F_PredicateLockPageSplit(m, l0, v995, v1015)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L61
	} else {
		goto L198
	}
L195:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1000+(v996^int32(-1))<<(uint(int32(6))%32))+16))
	v1015 = v1006
	goto L194
L196:
	;
	goto L197
L197:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1008+v996<<(uint(int32(6))%32)+int32(-64))+16))
	v1015 = v1014
	goto L194
L198:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v893)+28))
	if v1018 != 0 {
		v893 = v1018
		goto L174
	} else {
		goto L199
	}
L199:
	;
	goto L175
L200:
	;
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L61
	} else {
		goto L214
	}
L201:
	;
	if l5 == int32(1) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	goto L203
L203:
	;
	F_gistfillbuffer(m, v67, l4, l5, int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L61
	} else {
		goto L213
	}
L204:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1059)+6)))
	v1063 = F_PageIndexTupleOverwrite(m, v67, v7, v1059, v1060&int32(8191))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L61
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	F_PageIndexTupleDelete(m, v67, v7)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L61
	} else {
		goto L212
	}
L207:
	;
	if v1063 != 0 {
		goto L200
	} else {
		goto L208
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L61
	} else {
		goto L209
	}
L209:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1069 + int32(4)
	F_errmsg_internal(m, int32(670198), v29+int32(32))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L61
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(469697), int32(557), int32(387602))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L61
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	goto L203
L213:
	;
	goto L200
L214:
	;
	if l8 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L61
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	if l12 != 0 {
		v1118 = int64(1)
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L217
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = base.I64_rotr(v1118, int64(32))
	v1122 = int32(0)
	if l7 == v1122 {
		v2098 = v1122
		v2108 = v1118
		goto L6
	} else {
		goto L233
	}
L220:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094)+118)))
	if v1095 != int32(112) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1116 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L61
	} else {
		goto L232
	}
L222:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1099 <= int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1102 != 0 {
		goto L221
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	if base.Ui32(v1052&int32(65535)) <= base.Ui32(int32(2047)) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1103 != 0 {
		goto L221
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+48)) = uint16(v7)
	v1113 = int32(1)
	goto L230
L229:
	;
	v1113 = int32(0)
	goto L230
L230:
	;
	v1114 = F_gistXLogUpdate(m, l3, v29+int32(48), v1113, l4, l5, l8)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L61
	} else {
		goto L231
	}
L231:
	;
	v1118 = v1114
	goto L219
L232:
	;
	v1118 = v1116
	goto L219
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v49
	v2098 = v1122
	v2108 = v1118
	goto L6
L234:
	;
	F_errmsg_internal(m, int32(333445), int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L61
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(469697), int32(256), int32(387602))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L61
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(75)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v811
	F_errmsg_internal(m, int32(638920), v29)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L61
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(469697), int32(328), int32(387602))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L61
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	if l12 != 0 {
		goto L306
	} else {
		goto L307
	}
L241:
	;
	v1469 = v1428
	goto L274
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = l3
	if l3 < int32(0) {
		goto L257
	} else {
		goto L258
	}
L243:
	;
	v1193 = v752
	goto L246
L244:
	;
	goto L245
L245:
	;
	if v49 == int32(0) {
		goto L242
	} else {
		goto L255
	}
L246:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+16))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1193)))
	*(*int32)(unsafe.Add(mBase, uint32(v1206))) = base.I32_rotr(v1207, int32(16))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+16))
	v1212 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v1211)+4)) = uint16(v1212)
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+28))
	if v1214 != 0 {
		v1193 = v1214
		goto L246
	} else {
		goto L248
	}
L247:
	;
	if v49 == int32(0) {
		goto L242
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	v1230 = v752
	goto L250
L250:
	;
	v1244 = F_palloc(m, int32(8))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L61
	} else {
		goto L252
	}
L251:
	;
	v1428 = v752
	goto L241
L252:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1244))) = v1246
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+4)) = v1248
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v1251 = F_lappend(m, v1250, v1244)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L61
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v1251
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+28))
	if v1254 != 0 {
		v1230 = v1254
		goto L250
	} else {
		goto L254
	}
L254:
	;
	goto L251
L255:
	;
	v1624 = int32(0)
	v1632 = int32(1)
	goto L240
L256:
	;
	v1304 = F_PageGetTempPageCopySpecial(m, v1303)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L61
	} else {
		goto L260
	}
L257:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1289+(l3^int32(-1))<<(uint(int32(2))%32))))
	v1303 = v1295
	goto L256
L258:
	;
	goto L259
L259:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1303 = v1297 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L256
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1304
	v1307 = int32(0)
	v1308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1304)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1304+v1308)+12)) = uint16(v1307)
	if v752 != 0 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(0)
	v1414 = F_gistfillitupvec(m, v1386, v1398, v29+int32(60))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L61
	} else {
		goto L273
	}
L262:
	;
	v1325 = v752
	v1327 = v1307
	goto L265
L263:
	;
	goto L264
L264:
	;
	v1381 = F_palloc(m, int32(0))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L61
	} else {
		goto L272
	}
L265:
	;
	v1339 = v1327 + int32(1)
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1325)+28))
	if v1340 != 0 {
		v1325 = v1340
		v1327 = v1339
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v1344 = F_palloc(m, v1339<<(uint(int32(2))%32))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L61
	} else {
		goto L268
	}
L267:
	;
	goto L266
L268:
	;
	v1359 = int32(0)
	v1360 = v752
	goto L269
L269:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1344+v1359<<(uint(int32(2))%32)))) = v1375
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1360)+28))
	if v1379 != 0 {
		v1359 = v1359 + int32(1)
		v1360 = v1379
		goto L269
	} else {
		goto L271
	}
L270:
	;
	v1386 = v1344
	v1398 = v1339
	goto L261
L271:
	;
	goto L270
L272:
	;
	v1386 = v1381
	v1398 = v1307
	goto L261
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v752
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v1414
	v1428 = v29 + int32(48)
	goto L241
L274:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+4))
	if int32(0) < v1480 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	v1624 = v1428
	v1632 = v1592
	goto L240
L276:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+8))
	v1498 = v1483
	v1499 = int32(0)
	goto L279
L277:
	;
	goto L278
L278:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+28))
	if v1580 == int32(0) {
		v1587 = v874
		goto L293
	} else {
		goto L294
	}
L279:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+20))
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498)+6)))
	v1516 = v1499 + int32(1)
	v1520 = F_PageAddItemExtended(m, v1511, v1498, v1512&int32(8191), v1516&int32(65535), int32(0))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L61
	} else {
		goto L281
	}
L280:
	;
	goto L278
L281:
	;
	if v1520 == int32(0) {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	if l7 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498)+6)))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+4))
	if v1516 < v1552 {
		v1498 = v1498 + v1548&int32(8191)
		v1499 = v1516
		goto L279
	} else {
		goto L292
	}
L284:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498)+2)))
	v1528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498))))
	v1529 = int32(16)
	v1532 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526)+2)))
	v1533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526))))
	if v1527|v1528<<(uint(v1529)%32) == v1532|v1533<<(uint(v1529)%32) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	if v1543 == int32(0) {
		goto L283
	} else {
		goto L291
	}
L286:
	;
	goto L285
L287:
	;
	v1539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498)+4)))
	v1540 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526)+4)))
	if v1539 == v1540 {
		v1543 = int32(1)
		goto L286
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1543 = int32(0)
	goto L286
L290:
	;
	goto L289
L291:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1469)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v1546
	goto L283
L292:
	;
	goto L280
L293:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+20))
	v1589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1588)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1588+v1589)+8)) = v1587
	v1592 = int32(0)
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+20))
	v1594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1593)+16)))
	v1595 = v1593 + v1594
	v1596 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1595)+12)))
	if v11^int32(1) != 0 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1469)))
	if v1583 == int32(0) {
		v1587 = v874
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1580)))
	v1587 = v1586
	goto L293
L296:
	;
	v1601 = v1592
	goto L298
L297:
	;
	v1601 = int32(8)
	goto L298
L298:
	;
	if v49 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1603 = v1601
	goto L301
L300:
	;
	v1603 = int32(0)
	goto L301
L301:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+28))
	if v1605 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1606 = v1603
	goto L304
L303:
	;
	v1606 = int32(0)
	goto L304
L304:
	;
	v1607 = v1596&int32(65527) | v1606
	*(*uint16)(unsafe.Add(mBase, uint32(v1595)+12)) = uint16(v1607)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+20))
	v1610 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1609)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609+v1610))) = base.I32_wrap_i64(int64(base.Ui64(v876) >> (uint(int64(32)) % 64)))
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+20))
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1613)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v1613+v1614)+4)) = base.I32_wrap_i64(v876)
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+28))
	if v1617 != 0 {
		v1469 = v1617
		goto L274
	} else {
		goto L305
	}
L305:
	;
	goto L275
L306:
	;
	v1660 = int32(4437236)
	v1662 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1662 + int32(1)
	if v1632 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L307:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1644)+118)))
	if v1645 != int32(112) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1649 <= int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1652 != 0 {
		goto L306
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1654 = int32(1)
	F_XLogEnsureRecordSpace(m, v811, v811<<(uint(v1654)%32)|v1654)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L61
	} else {
		goto L314
	}
L312:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1653 != 0 {
		goto L306
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	goto L306
L315:
	;
	v1681 = v1624
	goto L318
L316:
	;
	goto L317
L317:
	;
	if l8 != 0 {
		goto L322
	} else {
		goto L323
	}
L318:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+24))
	F_MarkBufferDirty(m, v1694)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L61
	} else {
		goto L320
	}
L319:
	;
	goto L317
L320:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+28))
	if v1697 != 0 {
		v1681 = v1697
		goto L318
	} else {
		goto L321
	}
L321:
	;
	goto L319
L322:
	;
	F_MarkBufferDirty(m, l8)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L61
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+20))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+24))
	if v1727 < int32(0) {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	F_PageRestoreTempPage(m, v1726, v1745)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L61
	} else {
		goto L330
	}
L327:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1731+(v1727^int32(-1))<<(uint(int32(2))%32))))
	v1745 = v1737
	goto L326
L328:
	;
	goto L329
L329:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1745 = v1739 + v1727<<(uint(int32(13))%32) + int32(-8192)
	goto L326
L330:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+24))
	if v1748 < int32(0) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1624)+20)) = v1766
	if l12 != 0 {
		v1961 = int64(1)
		goto L335
	} else {
		goto L336
	}
L332:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1752+(v1748^int32(-1))<<(uint(int32(2))%32))))
	v1766 = v1758
	goto L331
L333:
	;
	goto L334
L334:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v1766 = v1760 + v1748<<(uint(int32(13))%32) + int32(-8192)
	goto L331
L335:
	;
	if v1632 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L336:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1769)+118)))
	if v1770 != int32(112) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1933 = F_gistGetFakeLSN(m, l0)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L61
	} else {
		goto L366
	}
L338:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1774 <= int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1777 != 0 {
		goto L337
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1779 = int32(0)
	v1780 = m.G0
	v1782 = v1780 - int32(32)
	m.G0 = v1782
	if v1624 != 0 {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1778 != 0 {
		goto L337
	} else {
		goto L343
	}
L343:
	;
	goto L341
L344:
	;
	v1785 = v1624
	v1786 = v1779
	goto L347
L345:
	;
	v1815 = v1779
	goto L346
L346:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1782)+28)) = uint8(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+26)) = uint16(v1815)
	*(*uint8)(unsafe.Add(mBase, uint32(v1782)+24)) = uint8(v815)
	*(*int64)(unsafe.Add(mBase, uint32(v1782)+16)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+8)) = v874
	F_XLogBeginInsert(m)
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L61
	} else {
		goto L350
	}
L347:
	;
	v1811 = v1786 + int32(1)
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+28))
	if v1812 != 0 {
		v1785 = v1812
		v1786 = v1811
		goto L347
	} else {
		goto L349
	}
L348:
	;
	v1815 = v1811
	goto L346
L349:
	;
	goto L348
L350:
	;
	if l8 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	F_XLogRegisterBuffer(m, int32(0), l8, int32(8))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L61
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	F_XLogRegisterData(m, v1782+int32(8), int32(24))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L61
	} else {
		goto L355
	}
L354:
	;
	goto L353
L355:
	;
	if v1624 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1856 = v1624
	v1858 = int32(1)
	goto L359
L357:
	;
	goto L358
L358:
	;
	v1928 = F_XLogInsert(m, int32(14), int32(48))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L61
	} else {
		goto L365
	}
L359:
	;
	v1883 = v1858 & int32(255)
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+24))
	F_XLogRegisterBuffer(m, v1883, v1884, int32(6))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L61
	} else {
		goto L361
	}
L360:
	;
	goto L358
L361:
	;
	v1888 = int32(4)
	F_XLogRegisterBufData(m, v1883, v1856+v1888, v1888)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L61
	} else {
		goto L362
	}
L362:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+8))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+12))
	F_XLogRegisterBufData(m, v1883, v1893, v1894)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L61
	} else {
		goto L363
	}
L363:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(v1856)+28))
	if v1899 != 0 {
		v1856 = v1899
		v1858 = v1858 + int32(1)
		goto L359
	} else {
		goto L364
	}
L364:
	;
	goto L360
L365:
	;
	m.G0 = v1782 + int32(32)
	v1961 = v1928
	goto L335
L366:
	;
	v1961 = v1933
	goto L335
L367:
	;
	v1981 = v1624
	goto L370
L368:
	;
	goto L369
L369:
	;
	if v49 != 0 {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1994)+4)) = base.I32_wrap_i64(v1961)
	*(*int32)(unsafe.Add(mBase, uint32(v1994))) = base.I32_wrap_i64(int64(base.Ui64(v1961) >> (uint(int64(32)) % 64)))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1981)+28))
	if v1997 != 0 {
		v1981 = v1997
		goto L370
	} else {
		goto L372
	}
L371:
	;
	goto L369
L372:
	;
	goto L371
L373:
	;
	v2098 = int32(1)
	v2108 = v1961
	goto L6
L374:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+28))
	if v2024 == int32(0) {
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v2040 = v2024
	goto L376
L376:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v2040)+24))
	F_UnlockReleaseBuffer(m, v2053)
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L61
	} else {
		goto L378
	}
L377:
	;
	goto L373
L378:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v2040)+28))
	if v2056 != 0 {
		v2040 = v2056
		goto L376
	} else {
		goto L379
	}
L379:
	;
	goto L377
L380:
	;
	if l8 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L381:
	;
	goto L382
L382:
	;
	v2150 = int32(4437236)
	v2152 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v2152 - int32(1)
	m.G0 = v29 + int32(864)
	return v2098
L383:
	;
	v2128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2127)+16)))
	v2132 = base.I32_wrap_i64(int64(base.Ui64(v2108) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v2128+v2127))) = v2132
	v2134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2127)+16)))
	v2136 = base.I32_wrap_i64(v2108)
	*(*int32)(unsafe.Add(mBase, uint32(v2127+v2134)+4)) = v2136
	v2138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2127)+16)))
	v2139 = v2127 + v2138
	v2140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2139)+12)))
	v2142 = v2140 & int32(65527)
	*(*uint16)(unsafe.Add(mBase, uint32(v2139)+12)) = uint16(v2142)
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+4)) = v2136
	*(*int32)(unsafe.Add(mBase, uint32(v2127))) = v2132
	goto L382
L384:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2113+(l8^int32(-1))<<(uint(int32(2))%32))))
	v2127 = v2119
	goto L383
L385:
	;
	goto L386
L386:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2127 = v2121 + l8<<(uint(int32(13))%32) + int32(-8192)
	goto L383
L387:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v2164 + int32(4)
	F_errmsg_internal(m, int32(670198), v29+int32(16))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L61
	} else {
		goto L388
	}
L388:
	;
	F_errfinish(m, int32(469697), int32(434), int32(387602))
	mBase = m.M
	v2177 = m.ExcPending
	if v2177 != 0 {
		goto L61
	} else {
		goto L389
	}
L389:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gtsvectorin(m *base.Module, l0 int32) int32 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(197700)
			F_errmsg(m, int32(182276), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(469501), int32(94), int32(261754))
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
