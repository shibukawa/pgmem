package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddWaitEventToSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	if l1 == int32(32) {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v10)
		v13 = int32(16)
	} else {
		v13 = l1
	}
	if l3 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_AddWaitEventToSet[0]))
		if v14 != v16 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_0), int32(0))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(587), int32(_a_F_AddWaitEventToSet_2))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v18 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_3), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(589), int32(_a_F_AddWaitEventToSet_2))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if v13&int32(1) != 0 {
					if v13&int32(134) != 0 {
						v41 = base.B2i32(l2 == int32(-1))
					} else {
						v41 = int32(0)
					}
					if v41 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_4), int32(0))
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(601), int32(_a_F_AddWaitEventToSet_2))
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = int32(1)
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 + v42
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v50 = v47 + v43<<(uint(int32(4))%32)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v13
						v58 = v13 - v42
						if v58 != 0 {
							if v58 == int32(15) {
								v65 = int32(_a_F_AddWaitEventToSet_5)
								v66 = v43
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
								*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v67
								v70 = v66
							} else {
								v70 = v43
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
							v65 = int32(_a_F_AddWaitEventToSet_6)
							v66 = v62
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
							*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v67
							v70 = v66
						}
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v74 = v71 + v70<<(uint(int32(3))%32)
						v75 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+6)) = uint16(v75)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v77
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						v81 = v79 - int32(1)
						if base.B2i32(v81 == v75)|base.B2i32(v81 == int32(15)) != 0 {
							v113 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v113)
						} else {
							v88 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v88)
							v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
							v91 = int32(1)
							v94 = int32(base.Ui32(v90)>>(uint(v91)%32)) & v91
							*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v94)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							if v96&int32(4) != 0 {
								v100 = v94 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v100)
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v103 = v100
								v104 = v102
							} else {
								v103 = v94
								v104 = v96
							}
							if v104&int32(128) == int32(0) {
							} else {
								v113 = v103 | int32(_a_F_AddWaitEventToSet_7)
								*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v113)
							}
						}
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_8), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(591), int32(_a_F_AddWaitEventToSet_2))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
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
	} else {
		if v13&int32(1) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_9), int32(0))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(596), int32(_a_F_AddWaitEventToSet_2))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if v13&int32(134) != 0 {
				v41 = base.B2i32(l2 == int32(-1))
			} else {
				v41 = int32(0)
			}
			if v41 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_AddWaitEventToSet_4), int32(0))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_AddWaitEventToSet_1), int32(601), int32(_a_F_AddWaitEventToSet_2))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v42 = int32(1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43 + v42
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v50 = v47 + v43<<(uint(int32(4))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v13
				v58 = v13 - v42
				if v58 != 0 {
					if v58 == int32(15) {
						v65 = int32(_a_F_AddWaitEventToSet_5)
						v66 = v43
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v67
						v70 = v66
					} else {
						v70 = v43
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62
					v65 = int32(_a_F_AddWaitEventToSet_6)
					v66 = v62
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v67
					v70 = v66
				}
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v74 = v71 + v70<<(uint(int32(3))%32)
				v75 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v74)+6)) = uint16(v75)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v74))) = v77
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				v81 = v79 - int32(1)
				if base.B2i32(v81 == v75)|base.B2i32(v81 == int32(15)) != 0 {
					v113 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v113)
				} else {
					v88 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v88)
					v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)))
					v91 = int32(1)
					v94 = int32(base.Ui32(v90)>>(uint(v91)%32)) & v91
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v94)
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
					if v96&int32(4) != 0 {
						v100 = v94 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v100)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						v103 = v100
						v104 = v102
					} else {
						v103 = v94
						v104 = v96
					}
					if v104&int32(128) == int32(0) {
					} else {
						v113 = v103 | int32(_a_F_AddWaitEventToSet_7)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v113)
					}
				}
				return
			}
		}
	}
}
func F_CreateWaitEventSet(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v6 = l1 << (uint(int32(4)) % 32)
	if l0 != 0 {
		F_ResourceOwnerEnlarge(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_CreateWaitEventSet[0]))
			v18 = F_MemoryContextAllocZero(m, v17, v6+l1<<(uint(int32(3))%32)+int32(32))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v20
				*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
				v26 = v18 + int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v26 + v6
				if l0 != 0 {
					F_ResourceOwnerRemember(m, l0, v18, int32(_a_F_CreateWaitEventSet_0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
						return v18
					}
				} else {
					return v18
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_CreateWaitEventSet[0]))
		v18 = F_MemoryContextAllocZero(m, v17, v6+l1<<(uint(int32(3))%32)+int32(32))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v20
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
			v26 = v18 + int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v26 + v6
			if l0 != 0 {
				F_ResourceOwnerRemember(m, l0, v18, int32(_a_F_CreateWaitEventSet_0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
					return v18
				}
			} else {
				return v18
			}
		}
	}
}
func F_EventTriggerAlterTableEnd(m *base.Module) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[0]))
	if v6 == int32(0) {
		return
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if v12 != 0 {
				v13 = int32(_a_F_EventTriggerAlterTableEnd_0)
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[1]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[1])) = v16
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
				v19 = F_lappend(m, v18, v10)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v19
					*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[1])) = v14
					v30 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v11
					return
				}
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableEnd[0]))
					v30 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v11
					return
				}
			}
		}
	}
}
func F_EventTriggerCollectSimpleCommand(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[0]))
	if v7 == int32(0) {
		return
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
		if v10 != 0 {
			return
		} else {
			v11 = int32(_a_F_EventTriggerCollectSimpleCommand_0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[1])) = v14
			v17 = F_palloc(m, int32(40))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(0)
				v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[2])))
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v22)
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+12)) = v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v26
				v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v30
				v32 = F_copyObjectImpl(m, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[0]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
					v38 = F_lappend(m, v37, v17)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v38
						*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCollectSimpleCommand[1])) = v12
						return
					}
				}
			}
		}
	}
}
func F_EventTriggerSQLDrop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(208)
	m.G0 = v10
	v15 = v2
	v16 = v2
	v17 = v2
	v18 = v2
	v19 = int32(-1)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v19 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v123 = int32(m.ExcTag)
	v124 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v123 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[0])) = v72
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[1])) = v71
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[2]))
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+8)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v73
	F_pg_re_throw(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L25
	}
L7:
	;
	m.G0 = v10 + int32(208)
	return
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[3])))
	if v23 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v71 = v15
	v72 = v16
	v73 = v17
	v74 = v18
	goto L10
L10:
	;
	if v74 != 0 {
		goto L6
	} else {
		goto L22
	}
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[4])))
	if v27&int32(1) == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[2]))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v17
	v46 = F_EventTriggerCommonSetup(m, l0, int32(2), int32(_a_F_EventTriggerSQLDrop_0), v10+int32(180))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	if v46 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v46
	F_CommandCounterIncrement(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[2]))
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+8)) = uint8(v58)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[0]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[1]))
	goto L18
L18:
	;
	v65 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v10 + int32(12)
	goto L21
L19:
	;
	v71 = v63
	v72 = v61
	v73 = v46
	v74 = int32(0)
	goto L10
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v72
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[1])) = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v73
	F_EventTriggerInvoke(m, v73, v10+int32(180))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[2]))
	v88 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+8)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[0])) = v72
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerSQLDrop[1])) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v73
	F_list_free(m, v73)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L7
L25:
	;
	goto L4
L26:
	;
	v128 = int32(v124)
	m.G0 = v10
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v10+int32(12) == v134 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	m.ExcPending = 1
	goto L35
L28:
	;
	if v138 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v138 = v136
	goto L31
L30:
	;
	v138 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+204))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v10)+200))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+196))
	v15 = v140
	v16 = v141
	v17 = v139
	v18 = v130
	v19 = v138
	goto L1
L33:
	;
	goto L34
L34:
	;
	F___wasm_longjmp(m, v131, v130)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_EventTriggerSupportsObjectType(m *base.Module, l0 int32) int32 {
	return (base.I32_wrap_i64(int64(base.Ui64(int64(4389322341887))>>(uint(base.I64_extend_i32_u(l0))%64))) | base.B2i32(base.Ui32(int32(42)) < base.Ui32(l0))) & int32(1)
}
func F_FreeWaitEventSet(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		F_ResourceOwnerForget(m, v3, l0, int32(_a_F_FreeWaitEventSet_0))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
