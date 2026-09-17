package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xactGetCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[0]))
	if v8 == v3 {
		v11 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v11)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
		return v11
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v17)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[1]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v38 = F_MemoryContextAlloc(m, v20, (v21+(v22+(v23+v24))-(v28+(v29+(v30+v31))))<<(uint(int32(4))%32))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v38
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[0]))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
			v47 = v45 - v46
			if int32(0) < v47 {
				v51 = v47 << (uint(int32(4)) % 32)
				if v51 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[2]))
					base.MemoryCopy(m, v38, v53+v46<<(uint(int32(4))%32), v51)
				} else {
				}
				v58 = v47
			} else {
				v58 = v3
			}
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
			v61 = v59 - v60
			if int32(0) < v61 {
				v65 = v61 << (uint(int32(4)) % 32)
				if v65 != 0 {
					v66 = int32(4)
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[2]))
					base.MemoryCopy(m, v38+v58<<(uint(v66)%32), v70+v60<<(uint(v66)%32), v65)
				} else {
				}
				v76 = v61 + v58
			} else {
				v76 = v58
			}
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
			v80 = v78 - v79
			if int32(0) < v80 {
				v84 = v80 << (uint(int32(4)) % 32)
				if v84 != 0 {
					v85 = int32(4)
					v89 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[3]))
					base.MemoryCopy(m, v38+v76<<(uint(v85)%32), v89+v79<<(uint(v85)%32), v84)
				} else {
				}
				v95 = v80 + v76
			} else {
				v95 = v76
			}
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
			v99 = v97 - v98
			if int32(0) < v99 {
				v103 = v99 << (uint(int32(4)) % 32)
				if v103 != 0 {
					v104 = int32(4)
					v108 = *(*int32)(unsafe.Add(mBase, _c_F_xactGetCommittedInvalidationMessages[3]))
					base.MemoryCopy(m, v38+v95<<(uint(v104)%32), v108+v98<<(uint(v104)%32), v103)
				} else {
				}
				v115 = v99 + v95
			} else {
				v115 = v95
			}
			return v115
		}
	}
}
func F_xact_redo_abort(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = v6 - int32(1)
	if v11 < int32(0) {
		v79 = l1
	} else {
		if v6&int32(1) != 0 {
			v16 = int32(2)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7+v11<<(uint(v16)%32))))
			if base.B2i32(base.Ui32(v16) < base.Ui32(v19))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l1)) == int32(0) {
				if base.Ui32(l1) < base.Ui32(v19) {
					v31 = v19
				} else {
					v31 = l1
				}
			} else {
				if int32(0) <= l1-v19 {
					v31 = l1
				} else {
					v31 = v19
				}
			}
			v34 = v31
			v36 = v6 - int32(2)
		} else {
			v34 = l1
			v36 = v11
		}
		if v11 == int32(0) {
			v79 = v34
		} else {
			v39 = v34
			v40 = v36
			for {
				v44 = int32(3)
				v48 = v7 + v40<<(uint(int32(2))%32)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				if base.B2i32(base.Ui32(v39) < base.Ui32(v44))|base.B2i32(base.Ui32(v49) < base.Ui32(v44)) == int32(0) {
					if v39-v49 < int32(0) {
						v59 = v49
					} else {
						v59 = v39
					}
				} else {
					if base.Ui32(v49) <= base.Ui32(v39) {
						v59 = v39
					} else {
						v59 = v49
					}
				}
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(4))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v62))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v59)) == int32(0) {
					if base.Ui32(v59) < base.Ui32(v62) {
						v74 = v62
					} else {
						v74 = v59
					}
				} else {
					if int32(0) <= v59-v62 {
						v74 = v59
					} else {
						v74 = v62
					}
				}
				if int32(1) < v40 {
					v39 = v74
					v40 = v40 - int32(2)
					continue
				} else {
					break
				}
				break
			}
			v79 = v74
		}
	}
	F_AdvanceNextFullTransactionIdPastXid(m, v79)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		return
	} else {
		v87 = *(*int32)(unsafe.Add(mBase, _c_F_xact_redo_abort[0]))
		if v87 == int32(0) {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_TransactionIdAbortTree(m, l1, v90, v91)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v113&int32(32) != 0 {
					v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
					v117 = int32(0)
					F_replorigin_advance(m, l3, v116, l2, v117, v117)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if int32(0) < v121 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								F_DropRelationFiles(m, v126, v127, int32(1))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v131 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v136, v137)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										return
									}
								}
							}
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if int32(0) < v131 {
								F_XLogFlush(m, l2)
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									F_pgstat_execute_transactional_drops(m, v136, v137)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								return
							}
						}
					}
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if int32(0) < v121 {
						F_XLogFlush(m, l2)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							F_DropRelationFiles(m, v126, v127, int32(1))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if int32(0) < v131 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										F_pgstat_execute_transactional_drops(m, v136, v137)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									return
								}
							}
						}
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if int32(0) < v131 {
							F_XLogFlush(m, l2)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_pgstat_execute_transactional_drops(m, v136, v137)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			F_RecordKnownAssignedTransactionIds(m, v79)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_TransactionIdAbortTree(m, l1, v96, v97)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_ExpireTreeKnownAssignedTransactionIds(m, l1, v100, v101, v79)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
						if v104&int32(64) == int32(0) {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
							if v113&int32(32) != 0 {
								v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
								v117 = int32(0)
								F_replorigin_advance(m, l3, v116, l2, v117, v117)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if int32(0) < v121 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_DropRelationFiles(m, v126, v127, int32(1))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v131 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v136, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													return
												}
											}
										}
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v131 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return
											} else {
												v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v136, v137)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return
												} else {
													return
												}
											}
										} else {
											return
										}
									}
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if int32(0) < v121 {
									F_XLogFlush(m, l2)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										F_DropRelationFiles(m, v126, v127, int32(1))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v131 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v136, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														return
													}
												}
											} else {
												return
											}
										}
									}
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if int32(0) < v131 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											F_pgstat_execute_transactional_drops(m, v136, v137)
											mBase = m.M
											v139 = m.ExcPending
											if v139 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										return
									}
								}
							}
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							F_StandbyReleaseLockTree(m, l1, v109, v110)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v113&int32(32) != 0 {
									v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
									v117 = int32(0)
									F_replorigin_advance(m, l3, v116, l2, v117, v117)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if int32(0) < v121 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												F_DropRelationFiles(m, v126, v127, int32(1))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if int32(0) < v131 {
														F_XLogFlush(m, l2)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return
														} else {
															v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															F_pgstat_execute_transactional_drops(m, v136, v137)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return
															} else {
																return
															}
														}
													} else {
														return
													}
												}
											}
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if int32(0) < v131 {
												F_XLogFlush(m, l2)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return
												} else {
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													F_pgstat_execute_transactional_drops(m, v136, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return
													} else {
														return
													}
												}
											} else {
												return
											}
										}
									}
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if int32(0) < v121 {
										F_XLogFlush(m, l2)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											F_DropRelationFiles(m, v126, v127, int32(1))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if int32(0) < v131 {
													F_XLogFlush(m, l2)
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return
													} else {
														v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														F_pgstat_execute_transactional_drops(m, v136, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return
														} else {
															return
														}
													}
												} else {
													return
												}
											}
										}
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if int32(0) < v131 {
											F_XLogFlush(m, l2)
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return
											} else {
												v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												F_pgstat_execute_transactional_drops(m, v136, v137)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return
												} else {
													return
												}
											}
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
