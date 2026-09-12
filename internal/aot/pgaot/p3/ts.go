package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ActivateCommitTs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v7 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v7 == int32(0) {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		v15 = F_LWLockAcquire(m, v11+int32(4992), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[85]))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
			v21 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			F_LWLockRelease(m, v21+int32(4992))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v19 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					v29 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					v32 = base.I32_div_u_s(v30, int32(819))
					v33 = base.I64_extend_i32_u(v32)
					*(*int64)(unsafe.Add(mBase, uint32(v27)+48)) = v33
					v36 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					v40 = F_LWLockAcquire(m, v36+int32(4992), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[87]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
						if v44 == int32(0) {
							v47 = F_ReadNextFullTransactionId(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, _consts[87]))
								v51 = base.I32_wrap_i64(v47)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v51
								v58 = *(*int32)(unsafe.Add(mBase, _consts[44]))
								F_LWLockRelease(m, v58+int32(4992))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									v64 = F_SimpleLruDoesPhysicalPageExist(m, int32(4365364), v33)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										if v64 == int32(0) {
											v69 = *(*int32)(unsafe.Add(mBase, _consts[86]))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
											v72 = int32(*(*uint16)(unsafe.Add(mBase, _consts[88])))
											v73 = base.I32_rem_u_s(v32, v72)
											v76 = v70 + v73<<(uint(int32(7))%32)
											v78 = F_LWLockAcquire(m, v76, int32(0))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v80 = int32(4365364)
												v82 = F_SimpleLruZeroPage(m, v80, v33)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return
												} else {
													F_SimpleLruWritePage(m, v80, v82)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														F_LWLockRelease(m, v76)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, _consts[44]))
															v94 = F_LWLockAcquire(m, v90+int32(4992), int32(0))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																v97 = *(*int32)(unsafe.Add(mBase, _consts[85]))
																v98 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
																v101 = *(*int32)(unsafe.Add(mBase, _consts[44]))
																F_LWLockRelease(m, v101+int32(4992))
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												}
											}
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, _consts[44]))
											v94 = F_LWLockAcquire(m, v90+int32(4992), int32(0))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _consts[85]))
												v98 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
												v101 = *(*int32)(unsafe.Add(mBase, _consts[44]))
												F_LWLockRelease(m, v101+int32(4992))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													return
												}
											}
										}
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, _consts[44]))
							F_LWLockRelease(m, v58+int32(4992))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = F_SimpleLruDoesPhysicalPageExist(m, int32(4365364), v33)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									if v64 == int32(0) {
										v69 = *(*int32)(unsafe.Add(mBase, _consts[86]))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
										v72 = int32(*(*uint16)(unsafe.Add(mBase, _consts[88])))
										v73 = base.I32_rem_u_s(v32, v72)
										v76 = v70 + v73<<(uint(int32(7))%32)
										v78 = F_LWLockAcquire(m, v76, int32(0))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = int32(4365364)
											v82 = F_SimpleLruZeroPage(m, v80, v33)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												F_SimpleLruWritePage(m, v80, v82)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													F_LWLockRelease(m, v76)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, _consts[44]))
														v94 = F_LWLockAcquire(m, v90+int32(4992), int32(0))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															v97 = *(*int32)(unsafe.Add(mBase, _consts[85]))
															v98 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
															v101 = *(*int32)(unsafe.Add(mBase, _consts[44]))
															F_LWLockRelease(m, v101+int32(4992))
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																return
															}
														}
													}
												}
											}
										}
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, _consts[44]))
										v94 = F_LWLockAcquire(m, v90+int32(4992), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _consts[85]))
											v98 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v98)
											v101 = *(*int32)(unsafe.Add(mBase, _consts[44]))
											F_LWLockRelease(m, v101+int32(4992))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
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
func F_TS_execute_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		if v17 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v20 == int32(1) {
					v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v92 = v24
						m.G0 = v10 + int32(16)
						return v92
					}
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					switch v26 - int32(1) {
					case 0:
						v29 = int32(1)
						if l2&v29 != 0 {
							v92 = v29
							m.G0 = v10 + int32(16)
							return v92
						} else {
							v34 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v34) < base.Ui32(int32(3)) {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[1048])))
									v92 = v86
								} else {
									v92 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v92
							}
						}
					case 1:
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v42 = F_TS_execute_recurse(m, l0+v38*int32(12), l1, l2, l3)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 == int32(0) {
								v92 = int32(0)
								m.G0 = v10 + int32(16)
								return v92
							} else {
								v48 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									switch v48 {
									case 0, 2:
										v92 = v48
									case 1:
										v92 = v42
									default:
										v92 = int32(0)
									}
									m.G0 = v10 + int32(16)
									return v92
								}
							}
						}
					case 2:
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v55 = F_TS_execute_recurse(m, l0+v51*int32(12), l1, l2, l3)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v55 == int32(1) {
								v92 = int32(1)
								m.G0 = v10 + int32(16)
								return v92
							} else {
								v61 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									switch v61 {
									case 0:
										v92 = v55
									case 1, 2:
										v92 = v61
									default:
										v92 = int32(0)
									}
									m.G0 = v10 + int32(16)
									return v92
								}
							}
						}
					case 3:
						v64 = F_TS_phrase_execute(m, l0, l1, l2, l3, int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							switch v64 {
							case 0, 1:
								v92 = v64
							case 2:
								v92 = l2 & int32(2)
							default:
								v92 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v92
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
							F_errmsg_internal(m, int32(474738), v10)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(488186), int32(1969), int32(355036))
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
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v20 == int32(1) {
				v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v92 = v24
					m.G0 = v10 + int32(16)
					return v92
				}
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				switch v26 - int32(1) {
				case 0:
					v29 = int32(1)
					if l2&v29 != 0 {
						v92 = v29
						m.G0 = v10 + int32(16)
						return v92
					} else {
						v34 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v34) < base.Ui32(int32(3)) {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[1048])))
								v92 = v86
							} else {
								v92 = int32(0)
							}
							m.G0 = v10 + int32(16)
							return v92
						}
					}
				case 1:
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v42 = F_TS_execute_recurse(m, l0+v38*int32(12), l1, l2, l3)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 == int32(0) {
							v92 = int32(0)
							m.G0 = v10 + int32(16)
							return v92
						} else {
							v48 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								switch v48 {
								case 0, 2:
									v92 = v48
								case 1:
									v92 = v42
								default:
									v92 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v92
							}
						}
					}
				case 2:
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v55 = F_TS_execute_recurse(m, l0+v51*int32(12), l1, l2, l3)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 == int32(1) {
							v92 = int32(1)
							m.G0 = v10 + int32(16)
							return v92
						} else {
							v61 = F_TS_execute_recurse(m, l0+int32(12), l1, l2, l3)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								switch v61 {
								case 0:
									v92 = v55
								case 1, 2:
									v92 = v61
								default:
									v92 = int32(0)
								}
								m.G0 = v10 + int32(16)
								return v92
							}
						}
					}
				case 3:
					v64 = F_TS_phrase_execute(m, l0, l1, l2, l3, int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						switch v64 {
						case 0, 1:
							v92 = v64
						case 2:
							v92 = l2 & int32(2)
						default:
							v92 = int32(0)
						}
						m.G0 = v10 + int32(16)
						return v92
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
						F_errmsg_internal(m, int32(474738), v10)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488186), int32(1969), int32(355036))
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
func F_TS_phrase_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		if v18 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v21 == int32(1) {
					v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, l4)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v281 = v24
						m.G0 = v11 + int32(48)
						return v281
					}
				} else {
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					switch v26 - int32(1) {
					case 0:
						v29 = int32(1)
						if l2&v29 != 0 {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v32)
							v281 = v29
							m.G0 = v11 + int32(48)
							return v281
						} else {
							v36 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, l4)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								switch v36 {
								case 0:
									v247 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
									v281 = v247
								case 1:
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									if int32(0) < v40 {
										v46 = (v39 ^ int32(-1)) & int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v46)
										v281 = int32(1)
									} else {
										if v39&int32(1) == int32(0) {
											v281 = int32(0)
										} else {
											v52 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v52)
											v281 = v52
										}
									}
								case 2:
									v281 = v36
								default:
									v281 = int32(0)
								}
								m.G0 = v11 + int32(48)
								return v281
							}
						}
					case 1, 3:
						v55 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v55
						*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v55
						*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v55
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v69 = F_TS_phrase_execute(m, l0+v63*int32(12), l1, l2, l3, v11+int32(32))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							if v69 == int32(0) {
								v281 = v6
								m.G0 = v11 + int32(48)
								return v281
							} else {
								v77 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v11+int32(16))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v77 == int32(0) {
										v281 = v6
										m.G0 = v11 + int32(48)
										return v281
									} else {
										v81 = int32(2)
										if v69 == v81 {
											v281 = v81
											m.G0 = v11 + int32(48)
											return v281
										} else {
											if v77 == int32(2) {
												v281 = v81
												m.G0 = v11 + int32(48)
												return v281
											} else {
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
												if v86 == int32(4) {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
													v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
													v91 = v89 + v90
													if l4 == int32(0) {
														v110 = v91
														v112 = int32(0)
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
														v106 = v91
														v107 = v91 + v95
														v108 = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v107
														v110 = v106
														v112 = v108
													}
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
													if v99 < v98 {
														v101 = v98
													} else {
														v101 = v99
													}
													v102 = v101 - v99
													v103 = v101 - v98
													if l4 == int32(0) {
														v110 = v103
														v112 = v102
													} else {
														v106 = v103
														v107 = v101
														v108 = v102
														*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v107
														v110 = v106
														v112 = v108
													}
												}
												v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
												v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
												if v114 == int32(1) {
													if v113&int32(1) != 0 {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
														v127 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(7), v110, v112, v124+v125)
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															if l4 != 0 {
																v247 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
																v281 = v247
															} else {
																v281 = int32(1)
															}
															m.G0 = v11 + int32(48)
															return v281
														}
													} else {
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														v136 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(2), v110, v112, v135)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															v281 = v136
															m.G0 = v11 + int32(48)
															return v281
														}
													}
												} else {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
													if v113&int32(1) != 0 {
														v146 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(1), v110, v112, v138)
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return int32(0)
														} else {
															v281 = v146
															m.G0 = v11 + int32(48)
															return v281
														}
													} else {
														v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
														if v138 < v153 {
															v155 = v138
														} else {
															v155 = v153
														}
														v156 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(4), v110, v112, v155)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															v281 = v156
															m.G0 = v11 + int32(48)
															return v281
														}
													}
												}
											}
										}
									}
								}
							}
						}
					case 2:
						v158 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v158
						*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v158
						*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v158
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v158
						v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v172 = F_TS_phrase_execute(m, l0+v166*int32(12), l1, l2, l3, v11+int32(32))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							v178 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v11+int32(16))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								if v172|v178 == int32(0) {
									v281 = int32(0)
									m.G0 = v11 + int32(48)
									return v281
								} else {
									v183 = int32(2)
									if v172 == v183 {
										v281 = v183
										m.G0 = v11 + int32(48)
										return v281
									} else {
										if v178 == int32(2) {
											v281 = v183
											m.G0 = v11 + int32(48)
											return v281
										} else {
											if v172 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(0)
											} else {
											}
											if v178 != 0 {
												v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												v196 = v192
											} else {
												v193 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v193
												v196 = v193
											}
											v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
											if v196 < v197 {
												v199 = v197
											} else {
												v199 = v196
											}
											*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v199
											v201 = v199 - v196
											v202 = v199 - v197
											v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
											v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
											if v204 == int32(1) {
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
												if v203&int32(1) != 0 {
													v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													if v207 < v215 {
														v217 = v207
													} else {
														v217 = v215
													}
													v218 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(4), v202, v201, v217)
													mBase = m.M
													v219 = m.ExcPending
													if v219 != 0 {
														return int32(0)
													} else {
														v247 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
														v281 = v247
														m.G0 = v11 + int32(48)
														return v281
													}
												} else {
													v220 = int32(1)
													v226 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), v220, v202, v201, v207)
													mBase = m.M
													v227 = m.ExcPending
													if v227 != 0 {
														return int32(0)
													} else {
														v228 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v228)
														v281 = v220
														m.G0 = v11 + int32(48)
														return v281
													}
												}
											} else {
												if v203&int32(1) == int32(0) {
													v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
													v258 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(7), v202, v201, v255+v256)
													mBase = m.M
													v259 = m.ExcPending
													if v259 != 0 {
														return int32(0)
													} else {
														v281 = v258
														m.G0 = v11 + int32(48)
														return v281
													}
												} else {
													v239 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v240 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(2), v202, v201, v239)
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return int32(0)
													} else {
														v247 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
														v281 = v247
														m.G0 = v11 + int32(48)
														return v281
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v263 = m.ExcPending
						if v263 != 0 {
							return int32(0)
						} else {
							v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v264
							F_errmsg_internal(m, int32(474738), v11)
							mBase = m.M
							v268 = m.ExcPending
							if v268 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(488186), int32(1837), int32(343351))
								mBase = m.M
								v273 = m.ExcPending
								if v273 != 0 {
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
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v21 == int32(1) {
				v24 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l0, l4)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v281 = v24
					m.G0 = v11 + int32(48)
					return v281
				}
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				switch v26 - int32(1) {
				case 0:
					v29 = int32(1)
					if l2&v29 != 0 {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v32)
						v281 = v29
						m.G0 = v11 + int32(48)
						return v281
					} else {
						v36 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, l4)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							switch v36 {
							case 0:
								v247 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
								v281 = v247
							case 1:
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if int32(0) < v40 {
									v46 = (v39 ^ int32(-1)) & int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v46)
									v281 = int32(1)
								} else {
									if v39&int32(1) == int32(0) {
										v281 = int32(0)
									} else {
										v52 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v52)
										v281 = v52
									}
								}
							case 2:
								v281 = v36
							default:
								v281 = int32(0)
							}
							m.G0 = v11 + int32(48)
							return v281
						}
					}
				case 1, 3:
					v55 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v55
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v55
					*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v55
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v55
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v69 = F_TS_phrase_execute(m, l0+v63*int32(12), l1, l2, l3, v11+int32(32))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						if v69 == int32(0) {
							v281 = v6
							m.G0 = v11 + int32(48)
							return v281
						} else {
							v77 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v11+int32(16))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								if v77 == int32(0) {
									v281 = v6
									m.G0 = v11 + int32(48)
									return v281
								} else {
									v81 = int32(2)
									if v69 == v81 {
										v281 = v81
										m.G0 = v11 + int32(48)
										return v281
									} else {
										if v77 == int32(2) {
											v281 = v81
											m.G0 = v11 + int32(48)
											return v281
										} else {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
											if v86 == int32(4) {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
												v91 = v89 + v90
												if l4 == int32(0) {
													v110 = v91
													v112 = int32(0)
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
													v106 = v91
													v107 = v91 + v95
													v108 = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v107
													v110 = v106
													v112 = v108
												}
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												if v99 < v98 {
													v101 = v98
												} else {
													v101 = v99
												}
												v102 = v101 - v99
												v103 = v101 - v98
												if l4 == int32(0) {
													v110 = v103
													v112 = v102
												} else {
													v106 = v103
													v107 = v101
													v108 = v102
													*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v107
													v110 = v106
													v112 = v108
												}
											}
											v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
											if v114 == int32(1) {
												if v113&int32(1) != 0 {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
													v127 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(7), v110, v112, v124+v125)
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int32(0)
													} else {
														if l4 != 0 {
															v247 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
															v281 = v247
														} else {
															v281 = int32(1)
														}
														m.G0 = v11 + int32(48)
														return v281
													}
												} else {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													v136 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(2), v110, v112, v135)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v281 = v136
														m.G0 = v11 + int32(48)
														return v281
													}
												}
											} else {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
												if v113&int32(1) != 0 {
													v146 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(1), v110, v112, v138)
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return int32(0)
													} else {
														v281 = v146
														m.G0 = v11 + int32(48)
														return v281
													}
												} else {
													v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
													if v138 < v153 {
														v155 = v138
													} else {
														v155 = v153
													}
													v156 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(4), v110, v112, v155)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														v281 = v156
														m.G0 = v11 + int32(48)
														return v281
													}
												}
											}
										}
									}
								}
							}
						}
					}
				case 2:
					v158 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v158
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v158
					*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v158
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v158
					v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v172 = F_TS_phrase_execute(m, l0+v166*int32(12), l1, l2, l3, v11+int32(32))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						v178 = F_TS_phrase_execute(m, l0+int32(12), l1, l2, l3, v11+int32(16))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int32(0)
						} else {
							if v172|v178 == int32(0) {
								v281 = int32(0)
								m.G0 = v11 + int32(48)
								return v281
							} else {
								v183 = int32(2)
								if v172 == v183 {
									v281 = v183
									m.G0 = v11 + int32(48)
									return v281
								} else {
									if v178 == int32(2) {
										v281 = v183
										m.G0 = v11 + int32(48)
										return v281
									} else {
										if v172 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(0)
										} else {
										}
										if v178 != 0 {
											v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											v196 = v192
										} else {
											v193 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v193
											v196 = v193
										}
										v197 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
										if v196 < v197 {
											v199 = v197
										} else {
											v199 = v196
										}
										*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v199
										v201 = v199 - v196
										v202 = v199 - v197
										v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
										v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
										if v204 == int32(1) {
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
											if v203&int32(1) != 0 {
												v215 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												if v207 < v215 {
													v217 = v207
												} else {
													v217 = v215
												}
												v218 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(4), v202, v201, v217)
												mBase = m.M
												v219 = m.ExcPending
												if v219 != 0 {
													return int32(0)
												} else {
													v247 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
													v281 = v247
													m.G0 = v11 + int32(48)
													return v281
												}
											} else {
												v220 = int32(1)
												v226 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), v220, v202, v201, v207)
												mBase = m.M
												v227 = m.ExcPending
												if v227 != 0 {
													return int32(0)
												} else {
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v228)
													v281 = v220
													m.G0 = v11 + int32(48)
													return v281
												}
											}
										} else {
											if v203&int32(1) == int32(0) {
												v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
												v258 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(7), v202, v201, v255+v256)
												mBase = m.M
												v259 = m.ExcPending
												if v259 != 0 {
													return int32(0)
												} else {
													v281 = v258
													m.G0 = v11 + int32(48)
													return v281
												}
											} else {
												v239 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
												v240 = F_TS_phrase_output(m, l4, v11+int32(32), v11+int32(16), int32(2), v202, v201, v239)
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return int32(0)
												} else {
													v247 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)) = uint8(v247)
													v281 = v247
													m.G0 = v11 + int32(48)
													return v281
												}
											}
										}
									}
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int32(0)
					} else {
						v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v264
						F_errmsg_internal(m, int32(474738), v11)
						mBase = m.M
						v268 = m.ExcPending
						if v268 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488186), int32(1837), int32(343351))
							mBase = m.M
							v273 = m.ExcPending
							if v273 != 0 {
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
func F_error_commit_ts_disabled(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_errcode(m, int32(325))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg(m, int32(496746), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
				if v20 == int32(1) {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+316))
					v28 = base.B2i32(v26 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v28)
					v30 = v28
				} else {
					v30 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(233336)
				if v30 != 0 {
					v35 = int32(584066)
				} else {
					v35 = int32(558292)
				}
				F_errhint(m, v35, v5)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errfinish(m, int32(486124), int32(390), int32(448309))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
func F_get_ts_parser_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(2281)
		v17 = int32(1)
		switch l1 - int32(5) {
		case 0:
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(9796820404457)
			v29 = int32(3)
			v31 = int32(2281)
			v32 = v29
		case 1:
			v31 = int32(2278)
			v32 = v17
		case 2:
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(15526306777321)
			v29 = int32(3)
			v31 = int32(2281)
			v32 = v29
		case 3:
			v29 = v17
			v31 = int32(2281)
			v32 = v29
		default:
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(23)
			v29 = int32(2)
			v31 = int32(2281)
			v32 = v29
		}
		v36 = F_LookupFuncName(m, v10, v32, v8+int32(20), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = F_get_func_rettype(m, v36)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				if v38 != v31 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v51 = F_func_signature_string(m, v10, v32, int32(0), v8+int32(20))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								v53 = F_format_type_be(m, v31)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v53
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51
									F_errmsg(m, int32(188892), v8)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486863), int32(126), int32(482370))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
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
				} else {
					m.G0 = v8 + int32(32)
					return v36
				}
			}
		}
	}
}
func F_ts_headline_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_DirectFunctionCall3Coll(m, int32(1189), int32(0), v4, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ts_headline_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_DirectFunctionCall3Coll(m, int32(1188), int32(0), v4, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ts_match_tq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	v2 = int32(0)
	v6 = l0 + int32(28)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_DirectFunctionCall1Coll(m, int32(1540), v2, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_DirectFunctionCall2Coll(m, int32(1538), v2, v17, v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if v23 != v7 {
						F_pfree(m, v7)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v19 != int32(0))
						}
					} else {
						return base.B2i32(v19 != int32(0))
					}
				}
			}
		}
	}
}
func F_ts_rank_tt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v16 = F_calc_rank(m, int32(1714480), v8, v14, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v18 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v16)
						}
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v22 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				} else {
					return base.I32_reinterpret_f32(v16)
				}
			}
		}
	}
}
func F_ts_rank_ttf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = F_calc_rank(m, int32(1714480), v8, v14, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v18 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v22 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v16)
						}
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v22 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return base.I32_reinterpret_f32(v16)
					}
				} else {
					return base.I32_reinterpret_f32(v16)
				}
			}
		}
	}
}
func F_ts_token_type_byname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = F_textToQualifiedNameList(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v18 = F_get_ts_parser_oid(m, v15, int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_tt_setup_firstcall(m, v13, l0, v18)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
							v25 = F_tt_process_call(m, v24)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								if v25 != 0 {
									v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
									*(*int64)(unsafe.Add(mBase, uint32(v24))) = v27 + int64(1)
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
									return v25
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
										v40 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
										return v25
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
		v25 = F_tt_process_call(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v25 != 0 {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
				*(*int64)(unsafe.Add(mBase, uint32(v24))) = v27 + int64(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
				return v25
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
					return v25
				}
			}
		}
	}
}
