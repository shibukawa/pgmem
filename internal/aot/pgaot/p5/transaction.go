package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsTransactionState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsTransactionState[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	return base.B2i32(v3 == int32(2))
}
func F_TransactionIdAbortTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	F_TransactionIdSetTreeStatus(m, l0, l1, l2, int32(2), int64(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_TransactionIdDidCommit(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[0]))
	if v10 == l0 {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[1]))
		v32 = v13
		v33 = int32(1)
		switch v32 - v33 {
		case 0:
			v75 = v33
			m.G0 = v7 + int32(16)
			return v75
		default:
			v75 = int32(0)
			m.G0 = v7 + int32(16)
			return v75
		case 2:
			v36 = int32(0)
			v37 = int32(3)
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[2]))
			if base.B2i32(base.Ui32(l0) < base.Ui32(v37))|base.B2i32(base.Ui32(v40) < base.Ui32(v37)) == v36 {
				if int32(0) <= l0-v40 {
					v50 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if v50 == int32(0) {
							v54 = int32(0)
							v57 = F_errstart(m, int32(19), v54)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 == int32(0) {
									v75 = v54
									m.G0 = v7 + int32(16)
									return v75
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										}
									}
								}
							}
						} else {
							v70 = F_TransactionIdDidCommit(m, v50)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v75 = v70
								m.G0 = v7 + int32(16)
								return v75
							}
						}
					}
				} else {
					v75 = v36
					m.G0 = v7 + int32(16)
					return v75
				}
			} else {
				if base.Ui32(l0) < base.Ui32(v40) {
					v75 = v36
					m.G0 = v7 + int32(16)
					return v75
				} else {
					v50 = F_SubTransGetParent(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if v50 == int32(0) {
							v54 = int32(0)
							v57 = F_errstart(m, int32(19), v54)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 == int32(0) {
									v75 = v54
									m.G0 = v7 + int32(16)
									return v75
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										}
									}
								}
							}
						} else {
							v70 = F_TransactionIdDidCommit(m, v50)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v75 = v70
								m.G0 = v7 + int32(16)
								return v75
							}
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			if l0 == int32(0) {
				v75 = int32(0)
			} else {
				v75 = int32(1)
			}
			m.G0 = v7 + int32(16)
			return v75
		} else {
			v21 = F_TransactionIdGetStatus(m, l0, v7+int32(8))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				switch v21 {
				case 0, 3:
					v32 = v21
				default:
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[1])) = v21
					*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[0])) = l0
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
					*(*int64)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[3])) = v30
					v32 = v21
				}
				v33 = int32(1)
				switch v32 - v33 {
				case 0:
					v75 = v33
					m.G0 = v7 + int32(16)
					return v75
				default:
					v75 = int32(0)
					m.G0 = v7 + int32(16)
					return v75
				case 2:
					v36 = int32(0)
					v37 = int32(3)
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdDidCommit[2]))
					if base.B2i32(base.Ui32(l0) < base.Ui32(v37))|base.B2i32(base.Ui32(v40) < base.Ui32(v37)) == v36 {
						if int32(0) <= l0-v40 {
							v50 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v54 = int32(0)
									v57 = F_errstart(m, int32(19), v54)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if v57 == int32(0) {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v75 = v54
													m.G0 = v7 + int32(16)
													return v75
												}
											}
										}
									}
								} else {
									v70 = F_TransactionIdDidCommit(m, v50)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = v70
										m.G0 = v7 + int32(16)
										return v75
									}
								}
							}
						} else {
							v75 = v36
							m.G0 = v7 + int32(16)
							return v75
						}
					} else {
						if base.Ui32(l0) < base.Ui32(v40) {
							v75 = v36
							m.G0 = v7 + int32(16)
							return v75
						} else {
							v50 = F_SubTransGetParent(m, l0)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v54 = int32(0)
									v57 = F_errstart(m, int32(19), v54)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if v57 == int32(0) {
											v75 = v54
											m.G0 = v7 + int32(16)
											return v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg_internal(m, int32(_a_F_TransactionIdDidCommit_0), v7)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_TransactionIdDidCommit_1), int32(162), int32(_a_F_TransactionIdDidCommit_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v75 = v54
													m.G0 = v7 + int32(16)
													return v75
												}
											}
										}
									}
								} else {
									v70 = F_TransactionIdDidCommit(m, v50)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = v70
										m.G0 = v7 + int32(16)
										return v75
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
func F_TransactionIdGetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
	v15 = base.I32_div_u_s(l0, int32(819))
	if l0 != 0 {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
			if l2 != 0 {
				v82 = v4
				v85 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v85)
				v125 = v82
			} else {
				v125 = v4
			}
			m.G0 = v11 + int32(16)
			return v125
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
			v25 = F_LWLockAcquire(m, v21+int32(_a_F_TransactionIdGetCommitTsData_0), int32(1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[1]))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
				if v31 == int32(0) {
					F_error_commit_ts_disabled(m)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					if l0 == v34 {
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = v36
						if l2 != 0 {
							v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v38)
						} else {
						}
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
						F_LWLockRelease(m, v41+int32(_a_F_TransactionIdGetCommitTsData_0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
							v125 = base.B2i32(v46 != int64(0))
							m.G0 = v11 + int32(16)
							return v125
						}
					} else {
						v49 = int32(0)
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[2]))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+44))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[0]))
						F_LWLockRelease(m, v55+int32(_a_F_TransactionIdGetCommitTsData_0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = int32(0)
							if base.B2i32(base.B2i32(v53 == v60)|base.B2i32(l0-v53 < v60)&base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))|base.B2i32(base.Ui32(v52) < base.Ui32(int32(3))) == v60)&base.B2i32(v60 <= v52-l0) != 0 {
								v91 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_TransactionIdGetCommitTsData_1), base.I64_extend_i32_u(v15), v11+int32(12))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[3]))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v91<<(uint(int32(2))%32))))
									v105 = v99 + (l0-v15*int32(819))*int32(10)
									v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+8)))
									v107 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
									*(*int64)(unsafe.Add(mBase, uint32(l1))) = v107
									if l2 != 0 {
										*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v106)
									} else {
									}
									v111 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[3]))
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
									v114 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdGetCommitTsData[4])))
									v115 = base.I32_rem_u_s(v15, v114)
									F_LWLockRelease(m, v112+v115<<(uint(int32(7))%32))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
										v125 = base.B2i32(v121 != int64(0))
										m.G0 = v11 + int32(16)
										return v125
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
								if l2 == int32(0) {
									v125 = v49
								} else {
									v82 = v49
									v85 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v85)
									v125 = v82
								}
								m.G0 = v11 + int32(16)
								return v125
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v135 = m.ExcPending
		if v135 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v138 = m.ExcPending
			if v138 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				F_errmsg(m, int32(_a_F_TransactionIdGetCommitTsData_2), v11)
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_TransactionIdGetCommitTsData_3), int32(296), int32(_a_F_TransactionIdGetCommitTsData_4))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
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
